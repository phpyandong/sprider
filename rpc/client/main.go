package main


import (
	"net"
	"net/rpc/jsonrpc"
	"sprider/rpc"
	"fmt"
	"log"
)

func main() {
	conn ,err := net.Dial("tcp","localhost:1234")
	if err != nil {
		panic("coon err")
	}
	client := jsonrpc.NewClient(conn)
	var result float64
	err = client.Call("DemoService.Div",rpcdemo.Args{11,33},&result)
	if err != nil {
		log.Printf("Demoservice.Div err %v",err)
	}
	fmt.Println(result)
}
