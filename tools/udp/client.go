// UDP 客户端
package main

import "net"
import (
	"fmt"
	"see/tools/udp/protocolV1"
	"time"
)

func sendData(data []byte) {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("连接服务端失败，err:", err)
		return
	}
	defer socket.Close()
	_, err = socket.Write(data) // 发送数据
	if err != nil {
		fmt.Println("发送数据失败，err:", err)
		return
	}
}
func report(moduleName, interfaceName, mes string) {

	p := &protocolV1.StatisticProtocol{
		ModuleNameLen: uint8(len(moduleName)),
		//接口名长度
		InterfaceNameLen: uint8(len(interfaceName)),
		//花费时间
		CostTime: float32(0.01),
		//是否成功
		Success: 1,
		//状态码
		Code: int32(200),
		//消息体的长度
		MsgLen: int16(len(mes)),
		//请求时间戳
		Time: int32(time.Now().Unix()),
		//模块名
		ModuleName: moduleName,
		//接口名
		InterfaceName: interfaceName,
		//消息体
		Msg: mes,
	}
	data, _ := p.ToBytes()
	sendData(data)
}

func main() {
	//
	//x := int32(123456)
	//bytesBuffer := bytes.NewBuffer([]byte{})
	//binary.Write(bytesBuffer, binary.BigEndian, x)
	//
	//m := md5.New()
	//m.Write(bytesBuffer.Bytes())
	//fmt.Println(hex.EncodeToString(m.Sum(nil)))

	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 55656,
	})
	if err != nil {
		fmt.Println("连接服务端失败，err:", err)
		return
	}
	defer socket.Close()
	/***
		方式一
	 */
	//p := new (protobuf.Protocol)
	//p.Format = []string{"C", "C", "f", "C", "N", "n", "N"}
	//aa := p.Pack(uint8(10),uint8(13),float32(5.5555),uint8(1),int32(391),int16(11),int32(time.Now().Unix()))
	//bb := []byte("TestModuleTestInterface我是message'")
	//ret := common.BytesCombine(aa,bb)
	/**
	方式二
	 */
	moduleName := "TestModule"
	interfaceName := "TestInterface"
	mes := "我是message"

	p := &protocolV1.StatisticProtocol{
		ModuleNameLen: uint8(len(moduleName)),
		//接口名长度
		InterfaceNameLen: uint8(len(interfaceName)),
		//花费时间
		CostTime: float32(2.2),
		//是否成功
		Success: 1,
		//状态码
		Code: int32(200),
		//消息体的长度
		MsgLen: int16(len(mes)),
		//请求时间戳
		Time: int32(time.Now().Unix()),
		//模块名
		ModuleName: moduleName,
		//接口名
		InterfaceName: interfaceName,
		//消息体
		Msg: mes,
	}

	//sendData := []byte("Hello server")
	mesBytes, _ := p.ToBytes()
	_, err = socket.Write(mesBytes) // 发送数据
	if err != nil {
		fmt.Println("发送数据失败，err:", err)
		return
	}
	data := make([]byte, 4096)
	n, remoteAddr, err := socket.ReadFromUDP(data) // 接收数据
	if err != nil {
		fmt.Println("接收数据失败，err:", err)
		return
	}
	fmt.Printf("recv:%v addr:%v count:%v\n", string(data[:n]), remoteAddr, n)
}
