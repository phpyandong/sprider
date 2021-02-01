package main

import (
	"sprider/craw/rpcsupport"
	"sprider/craw/store"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"flag"
	"os/signal"
	"syscall"
	"os"
	"net/http"
	"time"
	"context"
	"sprider/craw/config"
	"sprider/redis"
	"fmt"
)

var port = flag.Int("port", 0,
	"the port for me to listen on")
//go run main.go --port=1234

func main()  {
	//client, err := elastic.NewClient(
	//	elastic.SetSniff(false),
	//	elastic.SetURL("http://localhost:9200/"),
	//)
	//if err != nil {
	//	panic("client new err")
	//}
	//
	//rpcsupport.ServRpc(":123",store.ItemSaverService{
	//	Client:client,
	//})
	flag.Parse()
	if *port == 0 {
		*port = config.ItemSaverPost
	}
	host := fmt.Sprintf(":%d",*port)
	//初始化redis链接池
	redis.InitRedis()


	//time.Sleep(time.Second*5)
	//redis.GetRedis().Get()
	//time.Sleep(time.Second*5)
	//redis.GetRedis().Get()
	//time.Sleep(time.Second*5)
	//redis.GetRedis().Get()
	//time.Sleep(time.Second*5)
	//redis.GetRedis().Get()
	//time.Sleep(time.Second*5)
	//redis.GetRedis().Get()
	//log.Fatal(serverRpc(host,"data_profile"))
	log.Fatal(ServerGRpc(host,"data_profile"))

	//go HttpServer()


}
func HttpServer(){
	server := &http.Server{
		Addr:    ":8080",
	}
	http.HandleFunc("/ping", func(res http.ResponseWriter, req *http.Request){
		_, err := res.Write([]byte("pong"));
		if err != nil{
			log.Fatal("write err--->",err)
		}
	})
	server.ListenAndServe()
	// 用于捕获退出信号
	quit := make(chan os.Signal)

	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// 设置超时时间，两个心跳周期，假设一次心跳 3s
	ctx, cancelFunc := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancelFunc()

	// Shutdown 接口，如果没有新的连接了就会释放，传入超时 context
	// 调用这个接口会关闭服务，但是不会中断活动连接
	// 首先会将端口监听移除
	// 然后会关闭所有的空闲连接
	// 然后等待活动的连接变为空闲后关闭
	// 如果等待时间超过了传入的 context 的超时时间，就会强制退出
	// 调用这个接口 server 监听端口会返回 ErrServerClosed 错误
	// 注意，这个接口不会关闭和等待websocket这种被劫持的链接，如果做一些处理。可以使用 RegisterOnShutdown 注册一些清理的方法

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")
}
func ServerGRpc(host,index string) error{
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL("http://localhost:9200/"),
	)
	if err != nil {
		panic("client new err")
	}
	log.Printf("【%s】Es Client init:.... ：",rpcsupport.ProgramType)

	err = rpcsupport.ServGrpc(host,&store.ItemSaverService{
		Client:client,
		Index:index,
	})
	if err != nil {
		panic(err)
	}
	return err
}
func serverRpc(host,index string) error{
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL("http://localhost:9200/"),
	)
	if err != nil {
		panic("client new err")
	}

	err = rpcsupport.ServRpc(host,
	&store.ItemSaverService{//todo 注意这里要传引用，这样构成指针接收者？？？？
		Client:client,
		Index:index,
	},
	)
	if err != nil {
		panic(err)
	}
	return err
}
