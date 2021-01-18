package parser

import (
	"testing"
	"io/ioutil"
	"sprider/model"
	"sprider/core"
)

func TestParseProfile(t *testing.T) {
	contents,err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseProfile(contents,"http://album.zhenai.com/u/121212","健康快乐123")
	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element; but was %v",result.Items )
	}
	actual := result.Items[0]
	expected := core.Item{
		Url:"http://album.zhenai.com/u/121212",
		Type:"zhenai",
		Id :"121212",
		Payload:model.Profile{
			Age: 56,
			Name :"健康快乐123",
			Marrage:"离异",
			Height:168,
			Education:"大专",
			Occupation:"5001-8000元",
		},
	}
	t.Logf("expectd %v : \n was %v",expected,actual)

	if actual != expected {
		t.Errorf("expectd %v :but was %v",expected,actual)
	}
}
