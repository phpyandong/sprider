package sched

import (
	"sprider/core"
)
//engine 调用是利用channel 与schedule通信，替代函数调用
type CurrSched struct {
	requestChan chan core.Request //有人submit加入数据
	workChan chan chan core.Request  //chan core.Request 是worker 类型
	//createWork 中的 in 是哪里来的呢？是schd 选择了你这个channel以后，就会给你来发送数据

	//把一任务送给chan core.Request  是worker 类型 ，100个worker 100 个channal 灌入总channle
}

func (currSched *CurrSched) Submit(req core.Request) {
	currSched.requestChan <- req
}

func (this *CurrSched) ConfigureWorkChan(workquesChan chan core.Request) {
	//this.workChan = workquesChan
}
//有一个woker reday 了，可以接收request，替代配置workChan,方法调用
//会被无限次调用，不断地创建worker 创建后加入到队列中，保存，并调用workerReady 告知有可用的worker
func (this *CurrSched) WorkerReady(w chan core.Request) {
	this.workChan <- w
}

func (currSched *CurrSched) Run() {
	currSched.requestChan = make(chan core.Request)
	currSched.workChan = make(chan chan core.Request)
	go func() {
		//定义reuqest 队列 接收请求
		var requestQ []core.Request
		//定义worker 队列 接收worker
		var workQ []chan core.Request

		for {
			var activeRequ core.Request //当前请求没问题
			var activeWorker chan core.Request //todo why当前worker？？？
			//如果请求队列，和工作队列 均有数据的话，
			// 将数据[逐个]赋值给 当前 启用的 worker
			if len(requestQ) >0 && len(workQ) >0 {
				activeWorker = workQ[0]
				activeRequ = requestQ[0]
			}
			select {
				//读ques Chan加入到请求队列
				case r := <-currSched.requestChan:
					requestQ  = append(requestQ,r)
				//读work Chan加入到工作队列
				case w := <- currSched.workChan:
					workQ = append(workQ,w)
				//从当前活跃 request Chan 发送到 活跃的worker Chan
				case activeWorker <- activeRequ://默认为nil 不会select
					workQ = workQ[1:]//pop 出数据
					requestQ = requestQ[1:] //pop出数据

			}
		}
	}()
}
