package scheduler

import "sprider/engine"
//scheduler 实现 所有worker公用一个输入
type SeiyaScheduler struct {
	WorkerChan chan engine.Request//提交worker的channel

}

func (seiyasch *SeiyaScheduler) Submit(request engine.Request) {
	//把request 提交给 worker
	seiyasch.WorkerChan <- request
}

func (seiyasch *SeiyaScheduler) ConfigureMasterWorkerChan(re chan engine.Request) {
	seiyasch.WorkerChan = re
}

