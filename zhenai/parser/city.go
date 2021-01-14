package parser

import (
	"sprider/engine"
	"regexp"
)
//在城市详情页，获取用户列表，获取所有用户的个人主页列表
//const cityRe = `<a href="http://album.zhenai.com/u/([\d]+)" target="_blank">([^<]+)</a>`
const cityRe = `<a href="(http://album.zhenai.com/u/([\d]+))"([^>]+)>([^<]+)</a>`

//<a href="http://album.zhenai.com/u/1385132990" target="_blank">飞花落砚</a>
//func ParseCity(contents []byte) engine.ParseResult {
//	re := regexp.MustCompile(cityRe)
//	//[][][]byte
//	//newstr := `<a href="http://album.zhenai.com/u/1385132990" target="_blank">飞花落砚</a>`
//	//newstrBytes := []byte(newstr)
//	re.FindAllSubmatch(contents,-1)
//	//fmt.Println(string(contents),"matches")
//	//result := engine.ParseResult{}
//	//
//	return engine.NilParser(contents)
//}
func ParseCity(contents []byte) engine.ParseResult{
	//titles := regexp.MustCompile(`<title>([^<]+)</title>`)
	//mat := titles.FindAllSubmatch(contents,-1)
	//for _,v :=  range mat {
	//	fmt.Println(string(v[1]),"mat")
	//}

	re := regexp.MustCompile(cityRe)
	//[][][]byte
	//newstr := `<a href="http://album.zhenai.com/u/1385132990" target="_blank">飞花落砚</a>`
	//newstrBytes := []byte(newstr)
	matches := re.FindAllSubmatch(contents,-1)
	//fmt.Println(string(contents),"matches")
	result := engine.ParseResult{}
	limit :=100
	for _ ,m := range matches{
		if limit <= 0{
			break
		}
		name := string(m[4])

		result.Items = append(
			result.Items,"City Detail User " + name )
		result.Request = append(
			result.Request,
			engine.Request{
				Url: string(m[1]),
				ParserFunc: func(c []byte) engine.ParseResult{
					return ParseProfile(c,name)
				},
			},
		)
		limit--

	}
	return result
}