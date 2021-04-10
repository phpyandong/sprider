package main

import (
"context"
"fmt"
"log"
"math/rand"
"net/http"
"sync"
"time"

"golang.org/x/sync/errgroup"
)

const (
	TimeOutStatus   = 98
	ErrSystemStatus = 99
	Success         = 200
)

type SystemConfig struct {
	DefaultCount  int
	DefaultStep   int
	DefaultSecond int
	IsOpen        bool
}

// 函数验证
type Option func(*SystemConfig)

// 默认系统
var defaultConfig = &SystemConfig{DefaultCount: 10, DefaultStep: 5, DefaultSecond: 1}

// 设置是否开启窗口滑动
func SetIsOpen(isOpen bool) Option {
	return func(c *SystemConfig) {
		c.IsOpen = isOpen
	}
}

// 加载配置
func SetConfiguration(option ...Option) {
	for _, v := range option {
		v(defaultConfig)
	}
}

// 设置窗口长度和步长,exp:window(2,2)
func SetCountAndStep(count, step int) Option {
	return func(c *SystemConfig) {
		c.DefaultCount = count
		c.DefaultStep = step
	}
}

// 设置窗口时长（秒数）,exp:window(timestamp)
func SetSecond(second int) Option {
	return func(c *SystemConfig) {
		c.DefaultSecond = second
	}
}

type ErrSystemErr struct{}

func (e ErrSystemErr) Error() string {
	return fmt.Sprintf("%w%s", ErrSystemStatus, "system panic")
}

type TimeOutError struct{}

func (t TimeOutError) Error() string {
	return fmt.Sprintf("%w%s", TimeOutStatus, " time out")
}

type Call func(w http.Request, r *http.ResponseWriter) error

// 单个数据
type Element struct {
	creat  time.Time
	caller Call
	w      http.Request
	r      *http.ResponseWriter
	ch     chan struct{}
}

// 新建一个源

func newElement(dtNow time.Time, call Call) *Element {
	return &Element{creat: dtNow, caller: call, ch: make(chan struct{})}
}

func (e *Element) Call() error {
	if e.caller != nil {

		group, ctx := errgroup.WithContext(context.Background())
		group.Go(func() error {

			_, cancel := context.WithCancel(ctx)

			err := e.caller(e.w, e.r)

			if err == nil {
				e.ch <- struct{}{}
			} else {
				cancel()
				log.Println(err)
			}

			return err
		})

		// func的超时控制
		select {
		case <-e.ch:
			return nil
		case <-ctx.Done():
			close(e.ch)
			return ctx.Err()
		case <-time.After(1 * time.Second):
			return TimeOutError{}
		}
	}

	return ErrSystemErr{}
}

type ElementBucket struct {
	ElementMap   map[int64]*Element
	ElementSlice []*Element
	Length       int32
	mutex        sync.Mutex
	// 触发窗口计算信号
	sig chan struct{}
	// 上次统计时间
	lastTime time.Time
}
type ElementBuckets struct{
	ElementMap		map[int64]*Element
	ElementSlice	[]*Element
	Length			int32
	mutext 			sync.Mutex
	sig				chan struct{}
	lastTime		time.Time
}
func NewElementBucket() *ElementBucket {
	return &ElementBucket{ElementMap: make(map[int64]*Element, 0), ElementSlice: make([]*Element, 0, 10), Length: 0, lastTime: time.Now(), sig: make(chan struct{})}
}

// 加入窗口
func (e *ElementBucket) add(call Call, w http.Request, r *http.ResponseWriter) {
	//每来一个request请求，将请求时间，和请求体封装为一个请求体
	elemet := newElement(time.Now(), call)

	addFunc := func() int32 {
		e.mutex.Lock()
		defer e.mutex.Unlock()

		e.ElementMap[elemet.creat.UnixNano()] = elemet//将请求体加入到Bucket窗口中的按照时间方便操作删除 map中
		e.ElementSlice = append(e.ElementSlice, elemet)//将请求体加入到Bucket窗口中slice中
		e.Length++
		return e.Length
	}

	// 统计当前窗口的错误数量，也可以通过统计进入
	// 入口处的窗口数量，但是某个窗口期的错误数量/窗口时间频次应该更能为限流和降级做参考
	// 甚至可以统计错误类型，根据不同错误类型来采取不同的策略
	if err := elemet.Call(); err != nil {
		if length := addFunc(); length >= int32(defaultConfig.DefaultCount) {
			println("len:",length)
			e.sig <- struct{}{} //降级移动窗口
		}
	}
}

// 监控
func (e *ElementBucket) monitor() {

	for {
		select {
		case <-e.sig:
			e.moveWindow()
		case <-time.After(time.Duration(defaultConfig.DefaultSecond) * time.Second):
			e.moveTimeWindow(time.Now())
		}
	}
}

// 移动window
func (e *ElementBucket) moveWindow() {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	// 模拟上传，上报当前的流量计数。
	fmt.Printf("聚合一次数据,总长度：%v\n", len(e.ElementSlice))
	//看有多少个请求体。
	if len(e.ElementSlice) >= defaultConfig.DefaultStep {//长度大于默认步长5
		e.ElementSlice = e.ElementSlice[defaultConfig.DefaultStep:] //裁剪最新的步长
	}
	//两种方式，如果是错误信号的话。直接使用更换切片的方式，改变集合中的数据
	// 更换底层地址
	newSlice := make([]*Element, 0, len(e.ElementSlice))
	newSlice = append(newSlice, e.ElementSlice...)
	e.ElementSlice = newSlice
}

// 移动time window
func (e *ElementBucket) moveTimeWindow(dtNow time.Time) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	// 模拟上传
	fmt.Printf("聚合一次时间数据%v\n", len(e.ElementMap))
	//每隔一定时间 定时移动，时钟保证不停前进
	currentTime := e.lastTime.Add(time.Duration(defaultConfig.DefaultSecond) * time.Second)
	//上次更新时间加一段时间 得到当前时间
	currentTimeUnix := currentTime.UnixNano()
	//这就形成了窗口的移动
	for k := range e.ElementMap {
		//循环map窗口中的请求体,删除掉小于当前时间的请求体。保持时间区间内的记录，其他丢弃
		if k <= currentTimeUnix {
			delete(e.ElementMap, k)
		}
	}

	e.lastTime = currentTime //然后把最后更新时间置为当前时间
}

func main() {

	// 设定随机数
	rand.Seed(time.Now().UnixNano())

	bucket := NewElementBucket()//初始化窗口信息，上次修改时间等，初始化

	var wg sync.WaitGroup

	wg.Add(1)

	go bucket.monitor()//启动一个协程监控窗口移动

	for i := 0; i < 100; i++ {
		go func() {
			//模拟一百个请求，
			bucket.add(request, http.Request{}, nil)
		}()
	}

	wg.Wait()
}

// 模拟请求和执行时间
func request(w http.Request, r *http.ResponseWriter) error {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	return nil
}

