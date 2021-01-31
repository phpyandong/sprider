package rpcsupport

import (
	"net/rpc"
	"net"
	"net/rpc/jsonrpc"
	"log"
	pb "sprider/craw/rpcsupport/proto3"
	"google.golang.org/grpc"
	"fmt"
	"time"
)

const ProgramType = "SERVER"
func ServGrpc(host string,service pb.StoreServiceServer ) error{
	log.Printf("【%s】: ServGrpc host %s service :%v init:....",ProgramType,host,service)

	server := grpc.NewServer()
	pb.RegisterStoreServiceServer(server,service)
	listner, err := net.Listen("tcp",host)
	if err != nil {
		log.Fatalf("【%s】:failed to listen :%v",ProgramType,err)
	}
	log.Printf("【%s】:listing on %s",ProgramType,host)
	if err := server.Serve(listner); err != nil {
		log.Fatalf("【%s】: server Serve err %v",ProgramType,err)
	}
	log.Printf("【%s】:ServGrpc host %s service :%v ok:....",ProgramType,host,service)

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
			log.Printf("【%s】:accept err %v",ProgramType,err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
	return nil
}
func NewClient(host string)( *rpc.Client ,error){
	conn ,err := net.DialTimeout("tcp",host,time.Second)
	if err != nil {
		panic(err)
		return nil,err
	}
	client := jsonrpc.NewClient(conn)
	return client,nil
}
var GrpcConnPool []*grpc.ClientConn
func InitGrpcClient(host string)(error){
	log.Printf("InitGrpcClient host:%v: act...",host)

	for i:=0;i<1 ;i++  {
		//dialOption := grpc.WithReturnConnectionError()//开启连接失败返回错误，默认开启
		log.Printf("InitGrpcClient dial:%v: act...",host)

		conn,err := grpc.Dial(host,grpc.WithInsecure(),grpc.WithReturnConnectionError() )
		//todo 这里阻塞了，研究下
		log.Printf("InitGrpcClient dial:%v: end...",host)

		if err != nil {
			log.Printf("InitGrpcClient  err :host:%v:error :%v",host,err)
			//panic("gprc conn err")
		}else{
			if conn == nil {
				panic("conn is nil")
			}
			log.Printf("InitGrpcClient host:%v:ok",host)

			GrpcConnPool = append(GrpcConnPool,conn)
		}
	}
	log.Printf("InitGrpcClient host:%v: end...",host)

	log.Println(fmt.Sprintf("init grpc Len host :%v,%v",host,len(GrpcConnPool)))
	//panic("hhh")
	//if err != nil {
	//	//如果端口号不存在。
	//	//这里并不会报错，注意哦。todo
	//	log.Printf("Dial err %v",err)
	//	return nil ,err
	//}
	return nil
}
func NewGrpcClient(host string)(pb.StoreServiceClient,error) {
	conn := GetConn()
	client := pb.NewStoreServiceClient(conn)
	return client,nil

}
func GetConn() *grpc.ClientConn{
	if (len(GrpcConnPool) < 1) {
		log.Println("GetConn: err 获取连接失败")
	}
	fmt.Println("grpc len:",len(GrpcConnPool))
	conn := GrpcConnPool[0]
	GrpcConnPool = GrpcConnPool[1:]
	return conn
}

