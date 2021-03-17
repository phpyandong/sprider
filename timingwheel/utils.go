package timingwheel
import (
"sync"
"time"
)

// truncate returns the result of rounding x toward zero to a multiple of m.
//truncate返回将x向零舍入到m的倍数的结果。

// If m <= 0, Truncate returns x unchanged.
//每一层的 currentTime 都必须是 tickMs 的整数倍，
// 如果不满足则会将 currentTime 修剪为 tickMs 的整数倍。
// 修剪方法为：currentTime = startMs - (startMs % tickMs)；
//相当于舍去（非整数）小数部分
//相当于16秒要变为 15秒 （间隔3秒）的倍数
func truncate(x, m int64) int64 {
	if m <= 0 {
		//  2秒的间隔
		return x
	}
	// 16 - 16%3 =15
	return x - x%m
}

// timeToMs returns an integer number, which represents t in milliseconds.

//以毫秒为单位表示当前的时间
func timeToMs(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

// msToTime returns the UTC time corresponding to the given Unix time,
// t milliseconds since January 1, 1970 UTC.
func msToTime(t int64) time.Time {
	return time.Unix(0, t*int64(time.Millisecond)).UTC()
}

type waitGroupWrapper struct {
	sync.WaitGroup
}

func (w *waitGroupWrapper) Wrap(cb func()) {
	w.Add(1)
	go func() {
		cb()
		w.Done()
	}()
}
