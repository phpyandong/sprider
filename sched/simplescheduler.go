package sched

import (
	"sprider/core"
)


type Simplescheduler struct {
	WorkerChan chan core.Request
}

func (this *Simplescheduler) Submit(req core.Request) {
	this.WorkerChan <- req
}

func (this *Simplescheduler) ConfigCommWorkChan(req chan core.Request) {
	//使用方法调用配置woker Chan
	this.WorkerChan = req
}

