package tool

import "log"

/**
	防止野生goroutine
	由于main函数中的recovery 是无法recovery到 goroutinue 中的错误的，
	因此使用此函数创建协程，同时抑制panic ,记录日志，避免整个程序挂掉。
	创建协程使用此函数替代原生的go func()
 */
func Go(f func()){
	go func(){
		defer func(){
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
			}
		}()

		f()
	}()
}
//在一个http请求中不建议直接创建goroutinue
//使用如下模型，创建一个woker ,在http 请求中 创建一个message
type Message struct {

}
//将message 通过 channel 传递给 worker 去处理，避免了创建多个goroutinue
//详见爬虫worker
//chan <- &Message{}