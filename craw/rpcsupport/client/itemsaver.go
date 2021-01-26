package client

import (
	"sprider/craw/rpcsupport"
	"log"
	pb "sprider/craw/rpcsupport/proto3"
	"context"
)

const ProgramType = "Client"
func ItemStore(host string) (chan pb.Item,error){
	out := make(chan pb.Item)
	//client ,err := rpcsupport.NewClient(host)
	grpcClient ,err := rpcsupport.NewGrpcClient(host)
	if err != nil {
		return nil,err
	}



	go func() {
		itemCount := 0
		for   {

			item := <- out
			log.Printf("save items %v",item)
			log.Printf("【%s】save items：%v",ProgramType,item)


			res := &pb.SaveItemResult{}
			option := &pb.SaveItemRequest{Item: &item}
			res ,err = grpcClient.SaveItem(context.Background(),option)
			//err = client.Call("ItemSaverService.Save",
			//	item,&res)

			//err := Save(client,storeIndex,item)
			if err != nil {
				log.Printf("【%s】client:item saveStore err ,saveing item %v :%v",ProgramType,item,err)

			}
			log.Printf("【%s】client:item saveStore,saveing item %v res :%v",ProgramType,item,res)

			itemCount ++
		}
	}()
	return out,nil
}
