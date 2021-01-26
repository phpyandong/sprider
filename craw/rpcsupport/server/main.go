package main

import (
	"sprider/craw/rpcsupport"
	"sprider/craw/store"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"sprider/craw/config"
	"fmt"
)

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
	host := fmt.Sprintf(":%d",config.ItemSaverPost)

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
	log.Println("Es init:.... ：")

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
