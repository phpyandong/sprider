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
	log.Printf("%+v",err)

}

func getComment() error {
	commModel := new(model.Comment)
	id := int64(12121)
	err := commModel.SelectById(id)
	return errors.Wrap(err,"getComment")
}