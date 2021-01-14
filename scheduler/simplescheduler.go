package scheduler

import "sprider/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

//func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
//	panic("implement me")
//}
//
//func (s *SimpleScheduler) Run() {
//	panic("implement me")
//}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	//panic("implement me")
	s.workerChan = c
}

func (s *SimpleScheduler) Submit(request engine.Request){
	//send request down to worker chan
	s.workerChan <- request
}