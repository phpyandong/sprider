package main

import (
"fmt"
"net"
"os"
"os/signal"
"syscall"
"time"
	"reflect"
	"log"
)

const addr string = "127.0.0.1:8080"

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGUSR1, syscall.SIGUSR2)
	go server()
	//等待tcp server启动
	time.Sleep(2 * time.Second)
	client()
	fmt.Println("使用: ctrl+c 退出服务")
	<-c
	fmt.Println("服务退出")
}
var ConnPoolContainer  []net.Conn

func CreateConnPool() {

	conn1,err1 := net.Dial("tcp", "127.0.0.1:8080")
	if err1 == nil {
		ConnPoolContainer  = append(ConnPoolContainer,conn1)
	}
	conn2,err2 := net.Dial("tcp", "127.0.0.1:8080")
	if err2 == nil {
		ConnPoolContainer  = append(ConnPoolContainer,conn2)
	}else{
		panic(err2)
	}
	fmt.Println("len:",len(ConnPoolContainer))
}
func getConn() net.Conn{
	conn := ConnPoolContainer[0]
	ConnPoolContainer =  ConnPoolContainer[1:]

	return conn
}
func putConn(conn net.Conn){
	ConnPoolContainer = append(ConnPoolContainer,conn)
}
func client() {

	CreateConnPool()

	connClent := getConn()


	fmt.Printf("conn:类型%v",reflect.TypeOf(connClent))
	//connPool = append(connPool, connClent)
	len ,err := connClent.Write([] byte("哈哈哈"))
	if err != nil{
		log.Printf("conn1 err %v\n",err)
	}else{
		log.Printf("conn1 len %v\n",len)
	}
	connClent2 := getConn()
	putConn(connClent)

	connClent3 := getConn()
	len , err = connClent2.Write([] byte("哈哈哈2"))
	if err != nil{
		log.Printf("conn2 err %v",err)
	}else{
		log.Printf("conn2 len %v",len)
	}
	len , err = connClent3.Write([] byte("哈哈哈2"))
	if err != nil{
		log.Printf("conn2 err %v",err)
	}else{
		log.Printf("conn2 len %v",len)
	}
}

func server() {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("Error listening: ", err)
		os.Exit(1)
	}
	defer l.Close()
	fmt.Println("Listening on ", addr)
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err)
		}

		fmt.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())
		//go handleRequest(conn)
	}
}