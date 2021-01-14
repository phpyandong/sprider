package engine

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
	//WorkerReady(chan Request)
	//Run()
}