package fetcher

import (
	"net/http"
	"golang.org/x/text/transform"
	"io/ioutil"
	"fmt"
	"golang.org/x/text/encoding"
	"golang.org/x/net/html/charset"
	"bufio"
	"golang.org/x/text/encoding/unicode"
	"log"
	"regexp"
	"time"
)

var rateLimiter = time.Tick( time.Millisecond)
func Fetch(url string)([]byte,error){
	<-rateLimiter //控制速率，因为channel可以阻塞
	r ,ok := regexp.Compile("http://album.zhenai.com/u")
	if ok ==nil {
		match := r.MatchString(url)
		if match {
			return []byte(`<div class="des f-cl" data-v-4f6f1ada>阿坝 | 36岁 | 大专 | 离异 | 168cm | 5001-8000元</div>`),nil
		}

	}else{
	}
	resp,err := http.Get(url)
	if err != nil {
		panic(err)
		return nil ,err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		//fmt.Println("Error:status code",resp.StatusCode)
		return nil ,fmt.Errorf("Error:status code : %d",resp.StatusCode)
	}
	//将gbk 	转为utf8
	//newBody := transform.NewReader(resp.Body,simplifiedchinese.GBK.NewDecoder())
	//e := determinEncoding(resp.Body)
	//newBody := transform.NewReader(resp.Body,e.NewDecoder())

	//return  ioutil.ReadAll(newBody)
	bodyReader := bufio.NewReader(resp.Body)
	//获取编码类型
	e := determinEncoding(bodyReader)
	//转为utf-8
	utf8Reader := transform.NewReader(
		bodyReader,
		e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}
const htmlUser =`
 <div class="des f-cl" data-v-4f6f1ada>阿坝 | 36岁 | 大专 | 离异 | 168cm | 5001-8000元</div>
`
//读取内容的编码
//func determinEncoding(r io.Reader) encoding.Encoding{
//	bytes ,err := bufio.NewReader(r).Peek(1024)
func determinEncoding(r *bufio.Reader) encoding.Encoding{
		bytes ,err := r.Peek(1024)
	if err != nil{
		log.Printf("fetacher error %v",err)
		//panic(err)
		return unicode.UTF8
	}

	e,_,_ := charset.DetermineEncoding(
		bytes,"")
	return e
}