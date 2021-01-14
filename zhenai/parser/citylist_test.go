package parser

import (
	"testing"
	"io/ioutil"
)

func TestParseCityList(t *testing.T) {
	//contents ,err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	contents ,err := ioutil.ReadFile("city_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseCityList(contents)
	const resultSize = 1

	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		//"http://www.zhenai.com/zhenghun/akesu",
		//"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCities := []string{
		"阿坝",
		//"阿克苏","阿拉善盟",
	}

	if len(result.Request ) != resultSize{

		t.Errorf("result should have %d but has %d\n",resultSize,len(result.Request))
	}
	for i, url := range expectedUrls{
		if result.Request[i].Url != url {
			t.Errorf("expected url #%d: %s ;but was %s ",i,url,result.Request[i].Url)
		}
	}

	if len(result.Items) != resultSize {
		t.Errorf("items should have %d but has %d\n",resultSize,len(result.Items))
	}

	for i, city := range expectedCities{
		if result.Items[i].(string)!= city {
			t.Errorf("expected city #%d: %s ;but was %s ",i,city,result.Items[i].(string))
		}
	}

}