package core

import (
	"log"
	pb "sprider/craw/rpcsupport/proto3"
)

type CoreCurrEngine struct {
	Sched BaseCurrSched
	ItemChan chan pb.Item

}
type BaseCurrSched interface {
	Submit(Request)
	ConfigureWorkChan(chan Request)
	WorkerReady(chan Request)
	Run()
}

func (cu_engine CoreCurrEngine) Run(seeds ...Request){
	out := make(chan ParseResult)
	cu_engine.Sched.Run()
	for i:=0;i < 4;i++{
		go createWorker(out, cu_engine.Sched)
	}
	for _,r := range seeds{
		cu_engine.Sched.Submit(r)
	}
	for {
		res := <- out
		for _,item := range res.Items {
			log.Printf("cu_engine:line:29 Got item %v",item)
			go func() {
				cu_engine.ItemChan <- item
			}()
		}
		for _,request := range res.Request{
			go cu_engine.Sched.Submit(request)
		}
	}
}
func createWorker(out chan ParseResult,s BaseCurrSched){
	//创建一个请求channe用于通知 scheduler
	in := make(chan Request)
	for {
		//告知schedule 准备好了
		s.WorkerReady(in)
		requ := <- in
		result ,err := Worker(requ)
		if err != nil {
			continue
		}
		out <- result

	}
}