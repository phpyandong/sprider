package main

import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"strconv"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"log"
	"context"
)

// 模拟慢请求
func sleep(ctx *gin.Context) {
	t := ctx.Query("t")
	s, err := strconv.Atoi(t)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "参数错误: " + t})
		return
	}

	time.Sleep(time.Duration(s) * time.Second)
	ctx.JSON(http.StatusOK, gin.H{"msg": fmt.Sprintf("sleep %d s", s)})
}


const (
	stateHealth   = "health"
	stateUnHealth = "unhealth"
)

var state = stateHealth

//ab -n 10000 -c 200 "http://localhost:8099/health"
func health(ctx *gin.Context) {
	fmt.Println("health init ")

	status := http.StatusOK
	//if state == stateUnHealth {
	//	status = http.StatusServiceUnavailable
	//}
	Mss<- Message{timD:time.Now().Unix()}

	ctx.JSON(status, gin.H{"data": state})
}

//技术总结：
//
//1. Shutdown 方法要写在主 goroutine 中；
//
//2.在主 goroutine 中的处理逻辑才会阻塞等待处理；
//
//3.带超时的 Context 是在创建时就开始计时了，因此需要在接收到结束信号后再创建带超时的 Context。
//
//给大家推荐一个框架来快速构建带优雅退出功能的 http 服务，详见： https://www.cnblogs.com/zhucheer/p/12341595.html
//
func worker(i int,c chan Message){
	fmt.Println("woker created ")
	for
	{
		req  := <-c
		fmt.Printf("woker :%v get %v ",i,req)
	}

	return

}
type Message struct{
	timD int64
}
var Mss chan Message
func main() {
	done := make(chan error ,2)
	stop := make(chan struct{})
	go func() {
		done <- Serv(stop)
	}()
	go func() {
		done <- ServWork(stop)
	}()
	var stopped bool
	for i:=0; i< cap(done);i++ {
		if err := <-done; err != nil{
			fmt.Println("error:%v",err)
		}
		if !stopped {
			stopped = true
			close(stop)
		}
	}
}
func ServWork(stop <-chan struct{})  error{
	return nil
}
func Serv(stop chan struct{})  error{
	Mss = make(chan Message)
	for i:=0;i<10;i++{
		go worker(i,Mss)
	}
	e := gin.Default()
	e.GET("/health", health)
	e.GET("/sleep", sleep)

	server := &http.Server{
		Addr:    ":8099",
		Handler: e,
	}
	//+++++++++++++++++++++++++++++++++++++++++++++++
	//stop信号;另一个服务异常退出后，给出一个stop信号，关闭本服务
	go func() {
		<-stop
		server.Shutdown(context.Background())
	}()
	//+++++++++++++++++++++++++++++++++++++++++++++++
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server run err: %+v", err)
	}

	// 用于捕获退出信号
	quit := make(chan os.Signal)

	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-quit:
		// 捕获到退出信号之后将健康检查状态设置为 unhealth
		state = stateUnHealth
		log.Println("Shutting down state: ", state)
		// 设置超时时间，两个心跳周期，假设一次心跳 3s
		ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			log.Fatal("Server forced to shutdown:", err)
		}
	}
	log.Println("Shutting down server...")
	// Shutdown 接口，如果没有新的连接了就会释放，传入超时 context
	// 调用这个接口会关闭服务，但是不会中断活动连接
	// 首先会将端口监听移除
	// 然后会关闭所有的空闲连接
	// 然后等待活动的连接变为空闲后关闭
	// 如果等待时间超过了传入的 context 的超时时间，就会强制退出
	// 调用这个接口 server 监听端口会返回 ErrServerClosed 错误
	// 注意，这个接口不会关闭和等待websocket这种被劫持的链接，如果做一些处理。可以使用 RegisterOnShutdown 注册一些清理的方法
	log.Println("Server exiting")

	return nil
}
