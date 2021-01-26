package main

import (
	"testing"
	"sprider/craw/rpcsupport"
	"fmt"
	"sprider/craw/config"
	"time"
	"context"
	pb "sprider/craw/rpcsupport/proto3"
	"log"
)

//func TestItemSaver(t *testing.T){
//	host := fmt.Sprintf(":%d",config.ItemSaverPost)
//	//start ItemSaverServer
//	go serverRpc(host,"test1")
//	time.Sleep(time.Second)
//	client , err := rpcsupport.NewClient(host)
//	//start
//	item := core.Item{
//		Url:"www.bai",
//		Type :"zhenai",
//		Id :"100001",
//		Payload:model.Profile{
//			Name:"seiya",
//			Gender :"男",
//			Age: 29,
//			Height :180,
//		},
//	}
//	if err != nil {
//		panic("client_test err")
//	}
//	res := ""
//	err = client.Call("ItemSaverService.Save",
//		item,&res)
//	if err != nil || res != "ok" {
//		panic(err)
//		t.Errorf("result : %s;err :%s",res,err)
//	}
//}
func TestItem2Saver(t *testing.T){
	host := fmt.Sprintf(":%d",config.ItemSaverPost)
	//start ItemSaverServer
	go ServerGRpc(host,"test1")
	time.Sleep(time.Second)
	client , err := rpcsupport.NewGrpcClient(host)
	//start
	item := pb.Item{
		Url:  "www.bai",
		Type: "zhenai",
		Id:   "100001",
		Payload :nil,
		Car :nil,
	}

	if err != nil {
		panic("client_test err")
	}
	res := &pb.SaveItemResult{}
	option := &pb.SaveItemRequest{Item: &item}
	res ,err = client.SaveItem(context.Background(),option)

	if err != nil {
		panic(err)
		t.Errorf("result : %s;err :%s",res,err)
	}
	log.Printf("结果为：%v",res)
}
