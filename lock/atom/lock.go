package atom

import "sync/atomic"

type Spin int32
//自旋，直到拿到锁
func (s *Spin) Lock(){
	for !atomic.CompareAndSwapInt32((*int32)(s),0,1){

	}
}
func (s *Spin) Unlock(){
	atomic.StoreInt32( (*int32)(s),0)
}
