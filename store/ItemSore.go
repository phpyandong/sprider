package store

import (
	"log"
	"gopkg.in/olivere/elastic.v5"
	"context"
	"sprider/core"
	"errors"
)

func ItemStore() chan core.Item{
	out := make(chan core.Item)
	go func() {
		itemCount := 0
		for   {

			item := <- out
			log.Printf("save items %v",item)
			err := Save(item)
			if err != nil {
				log.Printf("item saveStore err ,saveing item %v :%v",item,err)

			}
			itemCount ++
		}
	}()
	return out
}

func Save(item core.Item)(err error){
	// Create a client and connect to http://192.168.2.10:9201
	//ES_HOST=es-hrnhpeom.public.tencentelasticsearch.com
	//ES_PORT=9200
	//ES_SCHEME=https
	//ES_USER=elastic
	//ES_PASS='eBEhGgmNb2So1gS#'
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL("http://localhost:9200/"),
		)
	if err != nil {
		//log.Printf("elactic connent err :%v",err)
		return err
	}
	if item.Type == "" {
		return errors.New("must supply Type")
	}
	indexService := client.Index().Index("test").
		Type(item.Type).
		Id(item.Id).
		BodyJson(item)
	if item.Id == "" {
		indexService.Id(item.Id)
	}
	 _, err = indexService.Do(context.Background())
	//fmt.Printf("%+v",result)
	if err != nil {
		return err
	}
	return nil
}

