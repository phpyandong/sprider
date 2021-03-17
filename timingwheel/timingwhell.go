package timingwheel

import (
	"unsafe"
	"container/list"
	"sync"
	"time"
	"sync/atomic"
)
//https://github.com/devYun/timingwheel/blob/master/timingwheel.go
//https://www.cnblogs.com/luozhiyun/p/14400363.html
type Timer struct {
	expiration int64 //任务过期时间
	task func()
	bucket unsafe.Pointer //type *bucket 所属的bucket
	element *list.Element //计时器的元素
}
//每多少秒为一个bucket
type bucket struct {
	expire int64 //任务过期时间
	mutex sync.Mutex
	timers *list.List //每个bucket的任务集合
}
func(b *bucket) Add(timer *Timer){
	b.mutex.Lock()
	b.timers.PushBack(timer)
	atomic.StorePointer(&timer.bucket,unsafe.Pointer(b))
	timer.bucket = b

}
func newBucket() *bucket{
	return &bucket{
		timers : list.New(),
		expire: -1,
	}
}
//时间轮由多个时间格组成，每个时间格代表当前时间轮的基本时间跨度（tickMs
//时间轮的时间格个数是固定的，可用 wheelSize 来表示
//那么整个时间轮的总体时间跨度（interval）可以通过公式 tickMs * wheelSize 计算得出。
type TimingWheel struct {
	//时间跨度 单位是秒
	tick int64
	//时间轮个数
	wheelSize int64
	//总跨度
	interval int64
	//当前指针指向的时间
	//时间轮还有一个表盘指针（currentTime），用来表示时间轮当前所处的时间，
	// currentTime 是 tickMs 的整数倍。
	// currentTime指向的地方是表示到期的时间格，表示需要处理的时间格所对应的链表中的所有任务。
	currentTime int64
	//时间格列表
	buckets []*bucket
	//上级时间轮引用
	overflowWheel unsafe.Pointer
	exitCh 	chan struct{}
	queue []interface{}

}
//3秒 4个格子
func NewTimingWheel(tick time.Duration,wheelSize int64) *TimingWheel{
	tickMs := int64(tick / 3*time.Second) //3秒一个格子
	startMs := time.Now().Unix() //当前时间戳 在这里转微秒
	return newTimingWheel(tickMs,wheelSize,startMs,make([]interface{},10))
}
//3s一个格子4个格子
func newTimingWheel (tickDuration int64,wheelSize int64, start int64,queue []interface{}) *TimingWheel{
	buckets := make([]*bucket,wheelSize)
	for i := range buckets{
		buckets[i] = newBucket()
	}
	return &TimingWheel{
		tick: tickDuration,//3秒
		wheelSize:wheelSize, //建立了一个轮子，他有4个格子
		interval:tickDuration * wheelSize,// 3 * 4 = 12 秒 //总格数
		//除了第一层时间轮，其余高层时间轮的起始时间（startMs）
		// 都设置为创建此层时间轮时前面第一轮的 currentTime。
		// 每一层的 currentTime 都必须是 tickMs 的整数倍，如果不满足则会将 currentTime 修剪为 tickMs 的整数倍。修剪方法为：currentTime = startMs - (startMs % tickMs)；
		currentTime:truncate(start,tickDuration),
		buckets : buckets,
		exitCh : make(chan struct{}),
		queue: queue,
	}

}
func (wheel *TimingWheel) add(t *Timer) bool{
	//拿到当前轮的当前时间
	currentTime := atomic.LoadInt64(&wheel.currentTime)
	if t.expiration < currentTime + wheel.tick{
		//new err 已经过期
		return false
	}else if t.expiration < currentTime/ wheel.interval{
		//小于12秒，第一个轮总跨度 11/3 = 3
		//获取时间轮的位置
		virtualID := t.expiration / wheel.tick //todo
		bucket := wheel.buckets[virtualID%wheel.wheelSize]//todo
		bucket.Add(t)
		if atomic.SwapInt64(&bucket.expire,virtualID * wheel.tick) != bucket.expire {
			wheel.queue = append(wheel.queue, bucket)//加入到队列
		}
		return  true
	}else{
		//如果放入的到期时间超过了第一层时间轮，那么就放到上一层中去
		overflowWheel := atomic.LoadPointer(&wheel.overflowWheel)
		if overflowWheel == nil {
			atomic.CompareAndSwapPointer(
				&wheel.overflowWheel,
				nil,
				//新的
				unsafe.Pointer(newTimingWheel(
					wheel.interval,  //第二圈的每块格子为12s，第一层的总跨度，第一层满12 进 1
					wheel.wheelSize,
					 currentTime,
					 wheel.queue,
					)),
			)

			overflowWheel = atomic.LoadPointer(&wheel.overflowWheel)

		}
		//往上递归
		return (*TimingWheel)(overflowWheel).add(t)
	}
}