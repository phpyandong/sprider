package parser

import (
	"testing"
	"io/ioutil"
	"sprider/model"
)

func TestParseProfile(t *testing.T) {
	contents,err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseProfile(contents,"健康快乐123")
	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element; but was %v",result.Items )
	}
	profile := result.Items[0].(model.Profile)
	expected := model.Profile{
		Age: 56,
		Name :"健康快乐123",
		Marrage:"离异",
		Height:168,
		Education:"大专",
		Occupation:"5001-8000元",
	}
	if profile != expected {
		t.Errorf("expectd %v :but was %v",expected,profile)
	}
}
