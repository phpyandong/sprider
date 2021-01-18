package engine

import (
	"log"
	"fmt"
)

type SeiyaEngine struct {
	Scheduler SeiyaBaseScheduler
	WorkerCount int  //worker 个数配置
}
func (seiyaengine *SeiyaEngine) Run(seeds ...Request){
	fmt.Println("run action .....")
	in := make(chan Request)
	out := make(chan ParseResult)
	//scheduler实现所有worker公用一个输入
	seiyaengine.Scheduler.ConfigureMasterWorkerChan(in)
	//创建5个worker ，通过两个channel 输入和输出进行数据的通信。
	for i := 0; i < seiyaengine.WorkerCount; i++ {
		go createSeiyaWorker(in,out)
	}
	//循环初始页。加入到调度任务中
	for _, r := range seeds{
		go seiyaengine.Scheduler.Submit(r)
	}

	log.Printf("deal the data from channel ...\n")
	//对channal 进行输入和输出处理
	itemCount := 0
	for {
		//获取输出的结果
		result := <- out
		//打印结果中的数据
		for _,item := range result.Items{
			log.Printf("Got item #%d:%v",itemCount,item)
			itemCount++
		}
		//循环将request 提交给worker
		for _,request := range result.Request {
			log.Printf("add to sch url:%v ",request.Url)

			go seiyaengine.Scheduler.Submit(request)
		}
	}
}

