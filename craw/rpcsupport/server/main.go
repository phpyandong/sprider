package main

import (
	"sprider/craw/rpcsupport"
	"sprider/craw/store"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"sprider/craw/config"
	"fmt"
	"sprider/redis"
	"flag"
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
