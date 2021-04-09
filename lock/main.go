package main

import (
	"flag"
	"sprider/lock/atom"
	"time"
	"sprider/lock/spin"
)

type Locker interface {
	Lock()
	Unlock()
}
func main(){
	var p bool
	flag.BoolVar(&p,"pause",false,"with pause")
	flag.Parse()

	var l Locker
	if p  {
		l = new(spin.Spin)
	}else{
		l = new(atom.Spin)

	}
	var n int
	for i:= 0;i<2 ;i++{
		go routinue(i,&n,l,500*time.Millisecond)
	}
	select{}
}

func routinue(i int,v *int,l Locker,d time.Duration){
	for  {
		//func(){ //加闭包的意思应该是让defer中的unlock 在方法体内生效，
		// 否则应该会在整个for循环外部生效，即不会释放锁，或者直接头部lock 尾部unlock
			l.Lock()
			//defer l.Unlock()
			*v++
			println(*v,i)
			time.Sleep(d)
		l.Unlock()
		//}()
	}
}
