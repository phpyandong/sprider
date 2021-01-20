package core

import (
	"sprider/fetcher"
	"log"
)

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