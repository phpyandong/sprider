package core

import (
	"log"
	"fmt"
	fetcher "sprider/fetcher"
)

type CoreEngine struct {
	Sched BaseSched
}
type BaseSched interface {
	Submit(request Request)
	ConfigCommWorkChan(chan Request)

}
func (coreEngine CoreEngine) Run(seeds ...Request)  {
	//
	fmt.Println("core engine action ...")
	in := make(chan Request)
	out := make(chan ParseResult)
	//需要配置公众输入channel
	coreEngine.Sched.ConfigCommWorkChan(in)
	//创建5个woker
	for i := 0; i < 4; i++ {
		fmt.Printf("create worker i:%v ...\n",i)

		go createW(in ,out)
	}
	for _,re := range seeds  {
		fmt.Printf("range seeds :%v ...\n",re)

		go coreEngine.Sched.Submit(re)
	}

	log.Printf("deal the data from channelx ...\n")

	for  {
		log.Printf("for begin ...\n")

		result := <-out
		log.Printf("get result from out :%v...\n",result)

		for _,item := range result.Items {
			log.Printf("输出的内容 %v",item)
		}
		for _,url := range result.Request{
			go coreEngine.Sched.Submit(url)
		}

	}

}
func createW(in chan Request,out chan ParseResult){
	for  {
		requ := <-in
		//in 是哪里来的呢？是schd 选择了你这个channel以后，就会给你来发送数据
		result,err:= Worker(requ)
		if err !=nil {
			fmt.Printf("err : createworker %v",err)
			continue
		}

		out <- result
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