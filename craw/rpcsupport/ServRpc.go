package rpcsupport

import (
	"net/rpc"
	"net"
	"net/rpc/jsonrpc"
	"log"
	pb "sprider/craw/rpcsupport/proto3"
	"google.golang.org/grpc"
)

func ServGrpc(host string,service pb.StoreServiceServer ) error{
	log.Printf("ServGrpc host %s service :%v init:....",host,service)

	server := grpc.NewServer()
	pb.RegisterStoreServiceServer(server,service)
	listner, err := net.Listen("tcp",host)
	if err != nil {
		log.Fatalf("failed to listen :%v",err)
	}
	log.Printf("listing on %s",host)
	if err := server.Serve(listner); err != nil {
		log.Fatalf("server Serve err %v",err)
	}
	log.Printf("ServGrpc host %s service :%v ok:....",host,service)

	return nil
}

func ServRpc(host string,service interface{}) error{
	//func ServRpc(host string,service *store.ItemSaverService) error{

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
func NewGrpcClient(host string)(pb.StoreServiceClient,error){

	conn,err := grpc.Dial(host,grpc.WithInsecure())
	if err != nil {
		return nil ,err
	}
	client := pb.NewStoreServiceClient(conn)
	return client ,nil
}
