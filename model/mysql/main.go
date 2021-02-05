package main

import (
	"sprider/config"
	"sprider/model"
	"github.com/pkg/errors"
	"log"
)

func main() {
	config.InitDB()

	err := getComment()
	log.Printf("堆栈信息%+v",err)
	log.Printf("%T %+v",errors.Cause(err),errors.Cause(err))


}

func getComment() error {
	commModel := new(model.Comment)
	id := int64(12121)
	err := commModel.SelectById(id)
	return errors.WithMessage(err,"getComment")
}