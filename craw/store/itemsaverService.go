package store

import (
	"gopkg.in/olivere/elastic.v5"
	"sprider/store"
	"log"
	pb "sprider/craw/rpcsupport/proto3"
	"context"
	"sprider/craw/rpcsupport"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index string
}

func (s *ItemSaverService) SaveItem(c context.Context,item *pb.SaveItemRequest) (*pb.SaveItemResult, error) {

	err := store.SaveGrpc(s.Client,s.Index,item.Item)
	if err == nil {
		log.Printf("【%s】Server:itemSaverService Save ok ：%v",rpcsupport.ProgramType,item)

	}else{
		log.Printf("【%s】Server:itemSaverService Save err ：%v",rpcsupport.ProgramType,err)
	}
	return &pb.SaveItemResult{},nil
}

func (s *ItemSaverService) Process(context.Context, *pb.ProcessRequest) (*pb.ProcessResult, error) {
	return nil,nil
}

//func (s *ItemSaverService) SaveItem(ctx context.Context,item core.Item,result *string) error{
//
//	err := store.SaveGrpc(s.Client,s.Index,item)
//	if err == nil {
//		*result = "ok"
//		log.Printf("Server:itemSaverService Save ok ：%v",item)
//
//	}else{
//		log.Printf("Server:itemSaverService Save err ：%v",err)
//	}
//	return err
//}
