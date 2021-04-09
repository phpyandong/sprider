package main

import (
	"sprider/pool/poolsample"
	"runtime"
	"fmt"
	"time"
)

func main(){
	sampool := poolsample.NewSamplePool(1)
	fmt.Println("goroutinue CHAN 3:",len(sampool.TaskQueue))

	w:= sampool.Get()
	w.Work()
	fmt.Println("goroutinue CHAN 1:",len(sampool.TaskQueue))

	fmt.Println("goroutinue Num 1:",runtime.NumGoroutine())

	time.Sleep(time.Second)
	sampool.Put(w)
	fmt.Println("goroutinue Num 2:",runtime.NumGoroutine())
	fmt.Println("goroutinue CHAN 2:",len(sampool.TaskQueue))


}

