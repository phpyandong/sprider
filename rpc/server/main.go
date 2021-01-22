package server

import (
	"net/rpc"
	"sprider/rpc"
	"net"
	"log"
	"net/rpc/jsonrpc"
)
// telnet localhost:1234
//{"method":"DemoService.Div","params":[{"A":3,"B":2}],"id":1}
func main()  {
	rpc.Register(rpcdemo.DemoService{})
	lister ,err  := net.Listen("tcp",":1234")
	if err != nil {
		panic("serve err")
	}

	for{
		conn,err := lister.Accept()
		if err != nil {
			log.Printf("accept err %v",err)
		}
		go jsonrpc.ServeConn(conn)
	}
}
