package store

import (
	"sprider/core"
	"gopkg.in/olivere/elastic.v5"
	"sprider/store"
	"log"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index string
}

func (s *ItemSaverService) Save(item core.Item,result *string) error{
	err := store.Save(s.Client,s.Index,item)
	if err == nil {
		*result = "ok"
	}else{
		log.Printf("Server:itemSaverService Save err ：%v",err)
	}
	return err
}
