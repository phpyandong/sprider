package model

import (
	"testing"
	"sprider/config"
	"fmt"
	"errors"
	"log"
	"sprider/basic"
)

func TestComment(t *testing.T) {
	config.InitDB()

	commModel := new(Comment)
	id := int64(12121)
	err := commModel.SelectById(id)
	if err != nil {
		if(errors.Is(err,basic.NotFoundError)){
			fmt.Println("404",basic.NotFoundError)
			return
		}
		log.Printf("err:%v",errors.Unwrap(err))
	}


}