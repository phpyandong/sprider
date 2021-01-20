package parser

import (
	"regexp"
	"strconv"
	"sprider/model"
	"fmt"
	"sprider/core"
	"strings"
)
var ageRegexp = regexp.MustCompile(`<td width="180"><span class="grayL">年龄：</span>([\d]+)</td>`)
var marrRegexp = regexp.MustCompile( `<td width="180"><span class="grayL">婚况：</span>([^<])</td>`)
var AllInfoRegexp = regexp.MustCompile(`des f-cl"([^>]+)>([^>]+)</div>`)
var idUrl  = regexp.MustCompile(
	`http://album.zhenai.com/u/([\d]+)`,
)
func ParseProfile(contents []byte,url string,name string) core.ParseResult {
	profile := model.Profile{}
	str := `<div class="des f-cl" data-v-4f6f1ada>阿坝 | 56岁 | 大专 | 离异 | 168cm | 5001-8000元</div>`
	strnew := []byte(str)
	fmt.Println(strnew)
	allInfoStr :=  extractString(contents,AllInfoRegexp)

	allInfo := strings.Split(allInfoStr, "|")

	profile.Name = strings.Trim(name," ")
	ageStr := strings.Trim(allInfo[1]," ")

	ageStr = strings.Trim(ageStr,"岁")
	ageint,err := strconv.Atoi(ageStr)
	if err != nil {
		ageint = 0
	}
	profile.Age = ageint
	profile.Education = strings.Trim(allInfo[2]," ")
	profile.Marrage = strings.Trim(allInfo[3]," ")
	Height := strings.Trim(allInfo[4]," ")

	Height = strings.Trim(Height,"cm")
	heightInt,err := strconv.Atoi(Height)
	if err != nil {
		heightInt = 0
	}
	profile.Height = heightInt
	profile.Occupation = strings.Trim(allInfo[5]," ")

	//result := core.ParseResult{
	//	//Items:[]interface{}{profile},
	//	Items:[]core.Item{
	//		{
	//			Url: url,
	//			Type :"zhenai",
	//			Id :extractStringUrl([]byte(url),idUrl),
	//			Payload:profile,
	//		},
	//
	//	},
	//}
	result := core.ParseResult{
		//Items:[]interface{}{profile},
		Items:[]core.Item{
			{
				Url: url,
				Type :"zhenai",
				Id :extractStringUrl([]byte(url),idUrl),
				Payload:profile,
			},

		},
	}
	return result
}

func ParseProfile2(contents []byte,name string) core.ParseResult{
	//re := regexp.MustCompile(ageRe)
	profile := model.Profile{}
	//用户名
	profile.Name = name
	//婚姻
	profile.Marrage = extractString(contents,marrRegexp)
	//年龄
	age ,err := strconv.Atoi(
		extractString(contents,ageRegexp),
	)
	fmt.Println(age,"年龄")
	if err != nil {
		profile.Age = age
	}
	//if match != nil {
	//	age,err := strconv.Atoi(string(match[1]))
	//	if err != nil {
	//		profile.Age = age
	//	}
	//
	//}
	//profile.Name
	result := core.ParseResult{
		Items:[]core.Item{
			{
				Url: "",
				Type :"zhenai",
				Id :"",
				Payload:profile,
			},

		},
	}
	return result
}
func extractStringUrl(contents []byte,re *regexp.Regexp ) string{
	match := re.FindSubmatch(contents)
	//fmt.Println(len(match),"match")
	//fmt.Println(string(match[0]),"match0")
	//fmt.Println(string(match[1]),"match1")
	//fmt.Println(string(match[2]),"match2")

	if len(match) >= 1 {
		return string(match[1])
	}
	return ""
}
func extractString(contents []byte,re *regexp.Regexp ) string{
	match := re.FindSubmatch(contents)
	//fmt.Println(len(match),"match")
	//fmt.Println(string(match[0]),"match0")
	//fmt.Println(string(match[1]),"match1")
	//fmt.Println(string(match[2]),"match2")

	if len(match) >= 2 {
		return string(match[2])
	}
	return ""
}