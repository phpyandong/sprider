package main


import (
	"sprider/core"
	"sprider/zhenai/parser"
	"sprider/sched"
	"github.com/pkg/errors"
	"sprider/craw/rpcsupport/client"
	"fmt"
	pb "sprider/craw/rpcsupport/proto3"
	"sprider/craw/rpcsupport"
	"log"
)



func main(){
	//2
	//e:= core.CoreEngine{
	//	Sched:&sched.Simplescheduler{},
	//
	//}
	//e.Run(core.Request{
	//	Url:"http://www.zhenai.com/zhenghun",
	//	ParserFunc:parser.ParseCityList,
	//})
	//e := engine.SeiyaEngine{
	//	Scheduler:&scheduler.SeiyaScheduler{},
	//	WorkerCount:3,
	//}
	//e.Run(engine.Request{
	//	Url:"http://www.zhenai.com/zhenghun",
	//	ParserFunc:parser.ParseCityList,
	//})
	//3
	//e := engine.ConcurrentEngine{
	//	Scheduler:&scheduler.QueuedScheduler{},
	//	WorkerCount:3,
	//}
	//e.Run(engine.Request{
	//	Url:"http://www.zhenai.com/zhenghun",
	//	ParserFunc:parser.ParseCityList,
	//})
	//itemChan ,err := client.ItemStore("localhost:1234") //store.ItemStore("data_profile")
	hosts := []string{":1234",":1235"}//todo 需要服务发现全部远程服务器地址，

	storeServiceClientChan ,err := createClientPool(hosts)
	if err != nil {
		panic(err)
	}
	itemChan ,err := client.ItemStoreChan(storeServiceClientChan) //store.ItemStore("data_profile")

	if err != nil {
		panic(
			errors.New(
				fmt.Sprintf("【%s】host localhost:1234 connect Err",
					client.ProgramType),
				),
			)
	}
	e := core.CoreCurrEngine{
		Sched:&sched.CurrSched{},
		ItemChan:itemChan,
	}
	e.Run(
		core.Request{
				Url:"http://www.zhenai.com/zhenghun",
				ParserFunc:parser.ParseCityList,
		},
	)
	//
	//// 1
	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:"http://www.zhenai.com/zhenghun",
	//	ParserFunc:parser.ParseCityList,
	//})
}
func createClientPool(hosts []string) (chan pb.StoreServiceClient,error){
	var clients []pb.StoreServiceClient
	for _ ,host :=  range hosts  {
		log.Printf("createClientPool hosts %v \n",host)

		client ,err := rpcsupport.NewGrpcClient(host)
		if err == nil {
			clients = append(clients,client)
		}else{
			panic("创建client err")
		}
	}
	out := make(chan pb.StoreServiceClient)
	go func() {
		for{
			for _ ,client :=  range clients  {
				log.Printf("createClientPool client %v",client)
				out <- client
			}
		}
	}()

	return out,nil
}

//
//func main2() {
//	resp,err := http.Get("http://www.zhenai.com/zhenghun")
//	if err != nil {
//		panic(err)
//	}
//
//	if resp.StatusCode != http.StatusOK {
//		fmt.Println("Error:status code",resp.StatusCode)
//		return
//	}
//	//将gbk 	转为utf8
//	//newBody := transform.NewReader(resp.Body,simplifiedchinese.GBK.NewDecoder())
//	e := determinEncoding(resp.Body)
//	newBody := transform.NewReader(resp.Body,e.NewDecoder())
//
//	all,err := ioutil.ReadAll(newBody)
//	if err != nil {
//		panic(err)
//	}
//	printCityList(all)
//	fmt.Sprintf("%s\n",all)
//}
////读取内容的编码
//func determinEncoding(r io.Reader) encoding.Encoding{
//	bytes ,err := bufio.NewReader(r).Peek(1024)
//	if err != nil{
//		panic(err)
//	}
//	e,_,_ := charset.DetermineEncoding(
//		bytes,"")
//	return e
//}
//func getRegex(){
//	text  := `m cc@gmailccc.com zhd
//m cc@gmailccc.org zhd
//m cc@qq.com zhd
//m cc@qq.com.cn zhd
//
//`
//	re := regexp.MustCompile(`([a-zA-Z]+)@([a-zA-Z0-9.]+)\.([a-zA-Z0-9]+)`)
//	//match := re.FindString(text)
//	//match := re.FindAllString(text,-1)
//	match := re.FindAllStringSubmatch(text,-1)
//	fmt.Println(match)
//}
//
//func printCityList(content []byte){
//	//< a data-v-1573aa7c href="http://www.zhenai.com/zhenghun/anhui">岳阳</a>
//	//<a data-v-1573aa7c="" href="http://www.zhenai.com/zhenghun/chuzhou">滁州</a>
//	re := regexp.MustCompile(
//		//`<a href="http://www.zhenai.com/zhenghun/[a-z0-9]+" data-v-1573aa7c>[^>]+</a>`,
//		`<a href="(http://www.zhenai.com/zhenghun/([a-z0-9]+))" data-v-1573aa7c>([^>]+)</a>`,
//	)
//	//[][][]byte
//	matches := re.FindAllSubmatch(content,-1)
//
//	for _ ,m := range matches{
//		fmt.Printf("City: %s , Url: %s\n ",m[3],m[1])
//
//		//for _, subMatch := range m{
//		//	fmt.Printf("City:%s Url:%s\n ",subMatch[2],subMatch[1])
//		//}
//		fmt.Println()
//		//fmt.Printf("%s\n",m)
//	}
//	fmt.Printf("mathch found : %d\n",len(matches))
//
//
//}