package rates

import (
"context"
"fmt"
"math"
"sync"
"time"
)

// Limit defines the maximum frequency of some events.
// Limit is represented as number of events per second.
// A zero Limit allows no events.
// Limit定义某些事件的最大频率。
// Limit表示为每秒事件数。
//零限制不允许任何事件。
type Limit float64

// Inf is the infinite rate limit; it allows all events (even if burst is zero).
// Inf为无限速率限制;它允许所有事件(即使burst为零)。
const Inf = Limit(math.MaxFloat64)

// Every converts a minimum time interval between events to a Limit.
// each将事件之间的最小时间间隔转换为一个限制。
func Every(interval time.Duration) Limit {
	if interval <= 0 {
		return Inf
	}
	return 1 / Limit(interval.Seconds())
}

// A Limiter controls how frequently events are allowed to happen.
// It implements a "token bucket" of size b, initially full and refilled
// at rate r tokens per second.
// Informally, in any large enough time interval, the Limiter limits the
// rate to r tokens per second, with a maximum burst size of b events.
// As a special case, if r == Inf (the infinite rate), b is ignored.
// See https://en.wikipedia.org/wiki/Token_bucket for more about token buckets.
//
// The zero value is a valid Limiter, but it will reject all events.
// Use NewLimiter to create non-zero Limiters.
//
// Limiter has three main methods, Allow, Reserve, and Wait.
// Most callers should use Wait.
//
// Each of the three methods consumes a single token.
// They differ in their behavior when no token is available.
// If no token is available, Allow returns false.
// If no token is available, Reserve returns a reservation for a future token
// and the amount of time the caller must wait before using it.
// If no token is available, Wait blocks until one can be obtained
// or its associated context.Context is canceled.
//
// The methods AllowN, ReserveN, and WaitN consume n tokens.
//限制器控制事件发生的频率。
//它实现了一个大小为b的“令牌桶”，最初是满的，然后重新填充
//速率r token / s。
//非正式地说，在任何足够大的时间间隔，限制器限制
//速率为每秒r个令牌，最大突发大小为b个事件。
//特殊情况下，如果r == Inf(无限速率)，b被忽略。
//参见https://en.wikipedia.org/wiki/Token_bucket了解更多关于令牌桶的信息。
//
// 0值是一个有效的限制器，但它将拒绝所有事件。
//使用newlimititer创建非零限制器。
//
//限制器有三个主要的方法，Allow, Reserve和Wait。
//大多数调用者应该使用Wait。
//
//三个方法中的每一个都消耗一个令牌。
//当没有可用的令牌时，它们的行为不同。
//如果没有可用的token, Allow返回false。
//如果没有可用的令牌，Reserve返回一个未来令牌的预留
//调用者在使用它之前必须等待的时间。
//如果没有可用的令牌，等待块直到一个可以获得
//或与其关联的上下文。上下文是取消了。
//
// allow, ReserveN和WaitN方法消耗n个token。
type Limiter struct {
	limit Limit
	burst int

	mu     sync.Mutex
	tokens float64
	// last is the last time the limiter's tokens field was updated
	// last是限制器的token字段最后一次更新
	last time.Time
	// lastEvent is the latest time of a rate-limited event (past or future)
	// lastEvent是限速事件的最新时间(过去或未来)
	lastEvent time.Time
}

// Limit returns the maximum overall event rate.
func (lim *Limiter) Limit() Limit {
	lim.mu.Lock()
	defer lim.mu.Unlock()
	return lim.limit
}

// Burst returns the maximum burst size. Burst is the maximum number of tokens
// that can be consumed in a single call to Allow, Reserve, or Wait, so higher
// Burst values allow more events to happen at once.
// A zero Burst allows no events, unless limit == Inf.
//Burst返回最大突发大小。Burst是允许、保留或等待的单个调用中可以消耗的最大令牌数，因此更高
//Burst值允许同时发生更多的事件。
//零爆发不允许任何事件，除非limit == Inf。
func (lim *Limiter) Burst() int {
	return lim.burst
}

// NewLimiter returns a new Limiter that allows events up to rate r and permits
// bursts of at most b tokens.
//newlimititer返回一个新的限制器，该限制器允许事件达到r级，并允许最多b个令牌的爆发。
func NewLimiter(r Limit, b int) *Limiter {
	return &Limiter{
		limit: r,
		burst: b,
	}
}

// Allow is shorthand for AllowN(time.Now(), 1).
//Allow是Allow (time.Now()， 1)的缩写。
func (lim *Limiter) Allow() bool {
	return lim.AllowN(time.Now(), 1)
}

// AllowN reports whether n events may happen at time now.
// Use this method if you intend to drop / skip events that exceed the rate limit.
// Otherwise use Reserve or Wait.
//allow报告当前是否可能发生n个事件。
//如果您想删除/跳过超过速率限制的事件，请使用此方法。
//否则使用预留或等待。
func (lim *Limiter) AllowN(now time.Time, n int) bool {
	return lim.reserveN(now, n, 0).ok
}

// A Reservation holds information about events that are permitted by a Limiter to happen after a delay.
// A Reservation may be canceled, which may enable the Limiter to permit additional events.
//一个Reservation保存了关于在延迟后被限制器允许发生的事件的信息。
//可以取消预订，这可能使限制器允许额外的事件。
type Reservation struct {
	ok        bool
	lim       *Limiter
	tokens    int
	timeToAct time.Time
	// This is the Limit at reservation time, it can change later.
	limit Limit
}

// OK returns whether the limiter can provide the requested number of tokens
// within the maximum wait time.  If OK is false, Delay returns InfDuration, and
// Cancel does nothing.
// OK返回限制器是否可以在最大等待时间内提供所请求的令牌数量。如果OK为false，则Delay返回InfDuration，而Cancel则不执行任何操作。
func (r *Reservation) OK() bool {
	return r.ok
}

// Delay is shorthand for DelayFrom(time.Now()).
func (r *Reservation) Delay() time.Duration {
	return r.DelayFrom(time.Now())
}

// InfDuration is the duration returned by Delay when a Reservation is not OK.
const InfDuration = time.Duration(1<<63 - 1)

// DelayFrom returns the duration for which the reservation holder must wait
// before taking the reserved action.  Zero duration means act immediately.
// InfDuration means the limiter cannot grant the tokens requested in this
// Reservation within the maximum wait time.
// DelayFrom返回保留操作的持有者在采取保留操作之前必须等待的时间。零持续时间意味着立即行动。
// induration表示限制器不能在最大等待时间内授予该预约中请求的令牌。
func (r *Reservation) DelayFrom(now time.Time) time.Duration {
	if !r.ok {
		return InfDuration
	}
	delay := r.timeToAct.Sub(now)
	if delay < 0 {
		return 0
	}
	return delay
}

// Cancel is shorthand for CancelAt(time.Now()).
// Cancel是CancelAt(time.Now())的简写。
func (r *Reservation) Cancel() {
	r.CancelAt(time.Now())
	return
}

// CancelAt indicates that the reservation holder will not perform the reserved action
// and reverses the effects of this Reservation on the rate limit as much as possible,
// considering that other reservations may have already been made.
//CancelAt表示预订持有者将不执行保留操作，并尽可能逆转该预订对速率限制的影响，
//考虑到可能已经提出了其他保留意见。
func (r *Reservation) CancelAt(now time.Time) {
	if !r.ok {
		return
	}

	r.lim.mu.Lock()
	defer r.lim.mu.Unlock()

	if r.lim.limit == Inf || r.tokens == 0 || r.timeToAct.Before(now) {
		return
	}

	// calculate tokens to restore
	// The duration between lim.lastEvent and r.timeToAct tells us how many tokens were reserved
	// after r was obtained. These tokens should not be restored.
	restoreTokens := float64(r.tokens) - r.limit.tokensFromDuration(r.lim.lastEvent.Sub(r.timeToAct))
	if restoreTokens <= 0 {
		return
	}
	// advance time to now
	now, _, tokens := r.lim.advance(now)
	// calculate new number of tokens
	tokens += restoreTokens
	if burst := float64(r.lim.burst); tokens > burst {
		tokens = burst
	}
	// update state
	r.lim.last = now
	r.lim.tokens = tokens
	if r.timeToAct == r.lim.lastEvent {
		prevEvent := r.timeToAct.Add(r.limit.durationFromTokens(float64(-r.tokens)))
		if !prevEvent.Before(now) {
			r.lim.lastEvent = prevEvent
		}
	}

	return
}

// Reserve is shorthand for ReserveN(time.Now(), 1).
/// Reserve是ReserveN(time.Now()， 1)的缩写。
func (lim *Limiter) Reserve() *Reservation {
	return lim.ReserveN(time.Now(), 1)
}

// ReserveN returns a Reservation that indicates how long the caller must wait before n events happen.
// The Limiter takes this Reservation into account when allowing future events.
// ReserveN returns false if n exceeds the Limiter's burst size.
//ReserveN返回一个预留，该预留指示调用者在n个事件发生之前必须等待多长时间。
//限制器在允许未来事件时考虑这个保留。
//如果n超过了限制器的突发大小，ReserveN返回false。
// Usage example:
//   r := lim.ReserveN(time.Now(), 1)
//   if !r.OK() {
//     // Not allowed to act! Did you remember to set lim.burst to be > 0 ?
//     return
//   }
//   time.Sleep(r.Delay())
//   Act()
// Use this method if you wish to wait and slow down in accordance with the rate limit without dropping events.
// If you need to respect a deadline or cancel the delay, use Wait instead.
// To drop or skip events exceeding rate limit, use Allow instead.
//如果您希望等待并按照速率限制放慢速度而不删除事件，请使用此方法。
//如果你需要尊重最后期限或取消延迟，使用Wait代替。
//要删除或跳过超过速率限制的事件，使用Allow代替。
func (lim *Limiter) ReserveN(now time.Time, n int) *Reservation {
	r := lim.reserveN(now, n, InfDuration)
	return &r
}

// Wait is shorthand for WaitN(ctx, 1).
func (lim *Limiter) Wait(ctx context.Context) (err error) {
	return lim.WaitN(ctx, 1)
}

// WaitN blocks until lim permits n events to happen.
// It returns an error if n exceeds the Limiter's burst size, the Context is
// canceled, or the expected wait time exceeds the Context's Deadline.
// The burst limit is ignored if the rate limit is Inf.
// WaitN阻塞直到lim允许n个事件发生。
//如果n超过了限制器的突发大小，上下文被取消，或者预期的等待时间超过了上下文的截止日期，则返回一个错误。
//如果速率限制为Inf，突发限制将被忽略。
func (lim *Limiter) WaitN(ctx context.Context, n int) (err error) {
	if n > lim.burst && lim.limit != Inf {
		return fmt.Errorf("rate: Wait(n=%d) exceeds limiter's burst %d", n, lim.burst)
	}
	// Check if ctx is already cancelled
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	// Determine wait limit
	now := time.Now()
	waitLimit := InfDuration
	if deadline, ok := ctx.Deadline(); ok {
		waitLimit = deadline.Sub(now)
	}
	// Reserve
	r := lim.reserveN(now, n, waitLimit)
	if !r.ok {
		return fmt.Errorf("rate: Wait(n=%d) would exceed context deadline", n)
	}
	// Wait if necessary
	delay := r.DelayFrom(now)
	if delay == 0 {
		return nil
	}
	t := time.NewTimer(delay)
	defer t.Stop()
	select {
	case <-t.C:
		// We can proceed.
		return nil
	case <-ctx.Done():
		// Context was canceled before we could proceed.  Cancel the
		// reservation, which may permit other events to proceed sooner.
		r.Cancel()
		return ctx.Err()
	}
}

// SetLimit is shorthand for SetLimitAt(time.Now(), newLimit).
//SetLimit是SetLimitAt(time.Now()， newLimit)的简写。
func (lim *Limiter) SetLimit(newLimit Limit) {
	lim.SetLimitAt(time.Now(), newLimit)
}

// SetLimitAt sets a new Limit for the limiter. The new Limit, and Burst, may be violated
// or underutilized by those which reserved (using Reserve or Wait) but did not yet act
// before SetLimitAt was called.
//SetLimitAt为限制器设置一个新的限制。新的限制和Burst可能被保留(使用Reserve或Wait)但在调用SetLimitAt之前尚未采取行动的那些人违反或未充分利用。
func (lim *Limiter) SetLimitAt(now time.Time, newLimit Limit) {
	lim.mu.Lock()
	defer lim.mu.Unlock()

	now, _, tokens := lim.advance(now)

	lim.last = now
	lim.tokens = tokens
	lim.limit = newLimit
}

// reserveN is a helper method for AllowN, ReserveN, and WaitN.
// maxFutureReserve specifies the maximum reservation wait duration allowed.
//// reserveN returns Reservation, not *Reservation, to avoid allocation in AllowN and WaitN.
//reserveN是allow、reserveN和WaitN的辅助方法。
// maxfuturerreserve允许的最大预留等待时间。
// reserveN返回预留，而不是*预留，以避免在allow和WaitN中分配。
func (lim *Limiter) reserveN(now time.Time, n int, maxFutureReserve time.Duration) Reservation {
	lim.mu.Lock()

	if lim.limit == Inf {
		lim.mu.Unlock()
		return Reservation{
			ok:        true,
			lim:       lim,
			tokens:    n,
			timeToAct: now,
		}
	}

	now, last, tokens := lim.advance(now)

	// Calculate the remaining number of tokens resulting from the request.
	//计算请求产生的token的剩余数量。
	tokens -= float64(n)
	fmt.Printf("token:%v\n",tokens)
	// Calculate the wait duration
	//计算等待时间
	var waitDuration time.Duration
	if tokens < 0 {
		waitDuration = lim.limit.durationFromTokens(-tokens)
	}

	// Decide result
	// 决定结果
	ok := n <= lim.burst && waitDuration <= maxFutureReserve

	// Prepare reservation
	//准备预订
	r := Reservation{
		ok:    ok,
		lim:   lim,
		limit: lim.limit,
	}
	if ok {
		r.tokens = n
		r.timeToAct = now.Add(waitDuration)
	}

	// Update state/ /更新状态
	if ok {
		lim.last = now
		lim.tokens = tokens
		lim.lastEvent = r.timeToAct
	} else {
		lim.last = last
	}

	lim.mu.Unlock()
	return r
}

// advance calculates and returns an updated state for lim resulting from the passage of time.
// lim is not changed.
// advance计算并返回lim随着时间的推移而更新的状态。
// lim不变。
func (lim *Limiter) advance(now time.Time) (newNow time.Time, newLast time.Time, newTokens float64) {
	last := lim.last
	if now.Before(last) {
		last = now
	}

	// Avoid making delta overflow below when last is very old.
	//当last非常老的时候，避免使下面的delta溢出。
	maxElapsed := lim.limit.durationFromTokens(float64(lim.burst) - lim.tokens)
	fmt.Printf("advance maxElapsed :%v \n",maxElapsed)
	elapsed := now.Sub(last)
	if elapsed > maxElapsed {
		elapsed = maxElapsed
	}

	// Calculate the new number of tokens, due to time that passed.
	//根据所经过的时间计算新的令牌数量。
	delta := lim.limit.tokensFromDuration(elapsed)
	fmt.Printf("advance增加：%v\n",delta)
	tokens := lim.tokens + delta
	if burst := float64(lim.burst); tokens > burst {
		tokens = burst
	}

	return now, last, tokens
}

// durationFromTokens is a unit conversion function from the number of tokens to the duration
// of time it takes to accumulate them at a rate of limit tokens per second.
//durationFromTokens是一个单位转换函数，从令牌数量到以每秒限制令牌的速度积累它们所需的时间。
func (limit Limit) durationFromTokens(tokens float64) time.Duration {
	fmt.Printf("durationFromTokens tokens :%v limit :%v\n",tokens,limit)
	seconds := tokens / float64(limit)
			//:= tokens总数/速率 = 20/2 =10s
	fmt.Printf("durationFromTokens tokens :%v limit :%v seconds: %v \n",tokens,limit,seconds)

	return time.Nanosecond * time.Duration(1e9*seconds)
}

// tokensFromDuration is a unit conversion function from a time duration to the number of tokens
// which could be accumulated during that duration at a rate of limit tokens per second.
func (limit Limit) tokensFromDuration(d time.Duration) float64 {
	return d.Seconds() * float64(limit)
}

