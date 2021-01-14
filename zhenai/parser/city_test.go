package parser

import (
	"testing"
	"io/ioutil"

	"fmt"
)

func TestParseCity(t *testing.T) {
	content ,err := ioutil.ReadFile("city_info.html")
	if err != nil {
		panic(err)
	}
	expectUser := "User 飞花落砚"
	userList := ParseCity(content)
	user := userList.Items[0]
	fmt.Println(user)
	if user != expectUser {
		panic(fmt.Sprintf("期待用户%s ，but %s",expectUser,user))
	}
	//fmt.Println(user)
	//fmt.Println(userList.Items)
	//fmt.Println(userList.Request)


}
