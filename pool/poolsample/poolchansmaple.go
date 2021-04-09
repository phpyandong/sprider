package poolsample

import "fmt"

type Worker interface {
	Work()
}

type EventWork struct {
	UrlId int
}
func(e *EventWork) Work(){
	fmt.Println(e.UrlId)
}
type SamplePool struct {
	TaskQueue chan Worker
}

func NewSamplePool(maxCoon int) *SamplePool  {
	samplePool := &SamplePool{
		TaskQueue:make(chan Worker,10),
	}
	for i:=0; i<maxCoon; i++ {
		wo := &EventWork{}
		samplePool.Put(wo)
	}
	return samplePool
}
func (s *SamplePool) Get() Worker{
	return <-s.TaskQueue
}

func(s *SamplePool) Put(worker Worker){
	s.TaskQueue<- worker
}
