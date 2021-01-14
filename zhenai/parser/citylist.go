package parser

import (
	"sprider/engine"
	"regexp"
)
const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/([a-z0-9]+))" data-v-1573aa7c>([^>]+)</a>`
//< a data-v-1573aa7c href="http://www.zhenai.com/zhenghun/anhui">岳阳</a>
//<a data-v-1573aa7c="" href="http://www.zhenai.com/zhenghun/chuzhou">滁州</a>
////`<a href="http://www.zhenai.com/zhenghun/[a-z0-9]+" data-v-1573aa7c>[^>]+</a>`,
//		`<a href="(http://www.zhenai.com/zhenghun/([a-z0-9]+))" data-v-1573aa7c>([^>]+)</a>`,
func ParseCityList(contents []byte) engine.ParseResult{
	re := regexp.MustCompile(cityListRe)
	//[][][]byte
	matches := re.FindAllSubmatch(contents,-1)
	result := engine.ParseResult{}
	limit :=10
	for _ ,m := range matches{
		if limit <=0 {
			break
		}
		result.Items = append(result.Items,"City"+string(m[3]))

		result.Request = append(
			result.Request,
			engine.Request{
				Url: string(m[1]),
				ParserFunc: func(cityRes []byte) engine.ParseResult {
					return ParseCity(cityRes)
				},
			},
		)
		limit--
	}
	return result
}