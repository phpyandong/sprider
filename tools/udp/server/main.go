package main

import (
	"net"
	"fmt"
)

// UDP server端
func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 55656,
	})

	// net.ListenTCP("tcp", &net.TCPAddr{
	//	IP:   net.IPv4(0, 0, 0, 0),
	//	Port: 55657,
	//})
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		var data [10240]byte
		n, addr, err := listen.ReadFromUDP(data[:]) // 接收数据

		fmt.Println("N:",n)
		if err != nil {
			fmt.Println("read udp failed, err:", err)
			continue
		}
		fmt.Println(data[0:n])
		/**
		方式一
		 */

		//p := new (protobuf.Protocol)
		//p.Format = []string{"C", "C", "f", "C", "N", "n", "N"}
		//fmt.Println(p.UnPack(data[0:n]))
		//fmt.Println(string(data[17:n]))
		/***
		方式二
		*/
		//p2 := new (protocolV1.StatisticProtocol)
		//p2.BytesToStruct(data[0:n])
		//fmt.Println(p2)
		//短整形16 n  占用2位
		//fmt.Println(uint(data[0]))
		//
		//fmt.Println(uint(data[0:n][0]))
		//
		//fmt.Println(int(binary.BigEndian.Uint16(data[0:2])))//


		//状态码N 占用4位 fmt.Println( int(binary.BigEndian.Uint32(data[0:4])))//

		//花费时间f  fmt.Println( math.Float32frombits(binary.LittleEndian.Uint32(data[0:4])))
		//aa := fmt.Sprintf("%.8f", math.Float32frombits(binary.BigEndian.Uint32(data[0:4])))

		//float_num, _ := strconv.ParseFloat(aa,32)
		//fmt.Println(float_num)


		//fmt.Printf("data:%v addr:%v count:%v\n",data[:n], addr, n)
		//fmt.Println(int(binary.BigEndian.Uint32(data[0:1])))
		//fmt.Println(int(binary.BigEndian.Uint32(data[1:2])))

		//
		//fmt.Println(data[0:n])
		//fmt.Println(string(data[0:n]))

		//fmt.Println(data[2:n])


		_, err = listen.WriteToUDP(data[:n], addr) // 发送数据
		if err != nil {
			fmt.Println("write to udp failed, err:", err)
			continue
		}
	}



}
