package client

import (
	"sprider/core"
	"sprider/craw/rpcsupport"
	"log"
)

func ItemStore(host string) (chan core.Item,error){
	out := make(chan core.Item)
	client ,err := rpcsupport.NewClient(host)
	if err != nil {
		return nil,err
	}



	go func() {
		itemCount := 0
		for   {

			item := <- out
			log.Printf("save items %v",item)
			res := ""
			err = client.Call("ItemSaverService.Save",
				item,&res)
			//err := Save(client,storeIndex,item)
			if err != nil {
				log.Printf("client:item saveStore err ,saveing item %v :%v",item,err)

			}
			itemCount ++
		}
	}()
	return out,nil
}
