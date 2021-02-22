package main
// 漏斗算法

import (
"fmt"
"math"
"sync"
"time"
)

func main() {
	service := &BucketLimiter{
		Timestamp: time.Now(),
		Capacity:  20,
		Rate:      2,
		Water:     0,
	}
	for true {
		hasToken := AddWater(service)
		hasToken2:= AddWater(service)

		if hasToken {
			fmt.Printf("hasToken:%v:token :%v:%v\n",time.Now(),hasToken,hasToken2)
		}

	}
}

type BucketLimiter struct {
	Timestamp time.Time // 当前注水的时间戳
	Capacity  float64   // 桶的容量
	Rate      float64   // 速度
	Water     float64   // 当前水量
	Lock      sync.Mutex
}

func AddWater(bucket *BucketLimiter) bool {
	now := time.Now()
	leftWater := math.Max(0, bucket.Water-now.Sub(bucket.Timestamp).Seconds()*bucket.Rate)
	bucket.Lock.Lock()
	defer bucket.Lock.Unlock()
	if leftWater+1 < bucket.Capacity {
		// 尝试加水，此时水桶未满
		bucket.Timestamp = now
		bucket.Water = leftWater + 1
		return true
	} else {
		// 水满了，拒绝访问
		return false
	}


}
