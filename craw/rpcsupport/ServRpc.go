package rpcsupport

import (
	"net/rpc"
	"net"
	"net/rpc/jsonrpc"
	"log"
)

func ServRpc(host string,service interface{}) error{
	rpc.Register(service)
	lister ,err  := net.Listen("tcp",host)
	if err != nil {
		//panic("serve err")
		return err
	}

	for{
		conn,err := lister.Accept()
		if err != nil {
			log.Printf("accept err %v",err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
	return nil
}
func NewClient(host string)( *rpc.Client ,error){
	conn ,err := net.Dial("tcp",host)
	if err != nil {
		panic(err)
		return nil,err
	}
	client := jsonrpc.NewClient(conn)
	return client,nil
}