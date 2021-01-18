package store

import (
	"testing"
	"sprider/model"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"context"
	"fmt"
	"encoding/json"
	"sprider/core"
)

//TestSave
func TestSave(t *testing.T){
	expected := core.Item{
		Url:"www.bai",
		Type :"zhenai",
		Id :"100001",
		Payload:model.Profile{
			Name:"seiya",
			Gender :"男",
			Age: 29,
			Height :180,
		},
	}
	err := Save(expected)
	fmt.Printf("id %s",expected.Id)
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL("http://localhost:9200/"),
	)
	if err != nil {
		log.Printf("elactic connent err :%v",err)
	}
	res ,err := client.Get().
		Index("test").
		Type(expected.Type).
		Id(expected.Id).Do(context.Background())
	fmt.Printf("res %+v",res.Source)
	//var actual model.Profile
	var actual core.Item
	err = json.Unmarshal(*res.Source,&actual)
	if err != nil {
		panic(err)
	}
	actualProfile ,_ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile
	if actual != expected{
		t.Errorf("got %v expect : %v",actual,expected)
	}

	//t.Errorf("got %v expect : %v",actual,user)


}