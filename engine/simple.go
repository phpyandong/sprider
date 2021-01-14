package engine

import (
	"sprider/fetcher"
	"log"
)
type SimpleEngine struct {}
func (e SimpleEngine) Run(seeds ...Request){
	var requests []Request
	for _, r := range seeds {
		requests = append(requests,r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("Fetching %s",r.Url)
		//获取指定url的body
		//body ,err := fetcher.Fetch(r.Url)
		//
		//if err != nil {
		//	log.Printf("Fetchcher:error  fetching url %s: %v",r.Url,err)
		//	continue
		//}
		////格式化body
		//ParseResult := r.ParserFunc(body)

		ParseResult,err := Worker(r)
		if err != nil {
			log.Printf("worker:error  %v",err)
		}


		requests = append(requests,ParseResult.Request...)

		//循环打印数据
		for _,item := range ParseResult.Items {
			log.Printf("Got Item %v",item)


		}
		//循环解析下一级

		for _,url := range ParseResult.Request {
			log.Printf("Got Url %v",url)

		}
	}

}

func Worker(r Request) (ParseResult,error){

	body,err := fetcher.Fetch(r.Url)
	log.Printf("FetchUrl: fetching"+
		"url %s ",r.Url)
	if err != nil {
		log.Printf("Fetch: error " + "fetching"+
			 "url %s : %v",r.Url,err)
		return ParseResult{},err
	}
	return r.ParserFunc(body),nil
}