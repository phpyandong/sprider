package engine

type BaseScheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
	WorkerReady(chan Request)
	Run()
}
type SeiyaBaseScheduler interface {
	Submit(Request)
	//scheduler实现所有worker公用一个输入 在这里配置这个传输的channel
	ConfigureMasterWorkerChan(chan Request)

}