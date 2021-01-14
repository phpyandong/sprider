package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}


func (e *ConcurrentEngine) Run(seeds ...Request){
	//for _, r := range seeds{
	//	e.Scheduler.Submit(r)
	//}
	//in := make(chan  Request)
	out := make(chan ParseResult)
	e.Scheduler.Run()

	//e.Scheduler.ConfigureMasterWorkerChan(in)

	for i := 0 ;i < e.WorkerCount ;i++ {
		//go createWorker(in ,out)
		go createWorker(out,e.Scheduler)

	}

	for _, r:= range seeds{
		e.Scheduler.Submit(r)
	}
	//j简单计数，看打印了多少数据
	itemCount :=0
	for {
		result := <- out
		for _,item := range result.Items{
			log.Printf("Got item #%d:%v",itemCount,item)
			itemCount++
		}
		for _,request := range result.Request {
			go e.Scheduler.Submit(request)
		}
	}
}
//func createWorker(in chan Request,out chan ParseResult)  {
func createWorker(out chan ParseResult,s Scheduler)  {
	in :=make(chan Request)
	//go func() {
		for {
			//todo tell schedule i'm ready
			s.WorkerReady(in)
			request := <- in
			result ,err := Worker(request)
			if err != nil {
				continue
			}
			out <-result
		}
	//}()
}

