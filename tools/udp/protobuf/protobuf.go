package protobuf

import (
	"math"
	"encoding/binary"
	"bytes"
	"see/tools/common"
	"see/tools/udp/protocolV1"

)

type Protocol struct{
	Format []string
}

func (p *Protocol) UnPack(msg []byte) ([]interface{}){

	la := len(p.Format)
	ret := make([]interface{}, la)
	if la > 0 {
		for i := 0; i < la; i++ {
			if p.Format[i] == "C" {
				ret[i] = uint(msg[0])
				msg = msg[1:len(msg)]

			}else if p.Format[i] == "f" {
				ret[i] = math.Float32frombits(binary.LittleEndian.Uint32(msg[0:4]))
				msg = msg[4:len(msg)]
			}else if p.Format[i] == "N" {
				ret[i] = int(binary.BigEndian.Uint32(msg[0:4]))
				msg = msg[4:len(msg)]
			}else if p.Format[i] == "n"{
				ret[i] = int(binary.BigEndian.Uint16(msg[0:2]))
				msg = msg[2:len(msg)]
			}
		}
	}
	return ret
}
func (p *Protocol) UnPackStruct(msg []byte) ([]interface{} ,*protocolV1.StatisticProtocol){
	sp := &protocolV1.StatisticProtocol{}
	buf := bytes.NewBuffer(msg)

	la := len(p.Format)
	ret := make([]interface{}, la)
	if la > 0 {
		for i := 0; i < la; i++ {
			if p.Format[i] == "C" {
				ret[i] = uint(msg[0])
				msg = msg[1:len(msg)]
				binary.Read(buf, binary.LittleEndian, &(sp.ModuleNameLen))

			}else if p.Format[i] == "f" {
				ret[i] = math.Float32frombits(binary.LittleEndian.Uint32(msg[0:4]))
				msg = msg[4:len(msg)]
			}else if p.Format[i] == "N" {
				ret[i] = int(binary.BigEndian.Uint32(msg[0:4]))
				msg = msg[4:len(msg)]
			}else if p.Format[i] == "n"{
				ret[i] = int(binary.BigEndian.Uint16(msg[0:2]))
				msg = msg[2:len(msg)]
			}
		}
	}
	return ret,sp
}
//封包
func (p *Protocol) Pack(args ...interface{}) []byte {
	la := len(args)
	ls := len(p.Format)
	ret := []byte{}
	if ls > 0 && la > 0 && ls == la {
		for i := 0; i < ls; i++ {
			if p.Format[i] == "C" {
				ret = append(ret,args[i].(uint8))
			}else if p.Format[i] == "N" {
				ret = append(ret, IntToBytes4(args[i].(int32))...)
			} else if p.Format[i] == "n" {
				ret = append(ret, IntToBytes2(args[i].(int16))...)
			} else  if p.Format[i] == "f"  {
				by := make([]byte,4)
				nowValue := math.Float32bits(args[i].(float32))
				binary.LittleEndian.PutUint32(by,nowValue)
				//合并两个bytes
				ret = common.BytesCombine(ret,by)
			}
		}
	}
	return ret
}


//整形转换成字节4位
func IntToBytes4(n int32) []byte {
	//m := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.LittleEndian, n)

	gbyte := bytesBuffer.Bytes()
	//c++ 高低位转换
	k := 4
	x := len(gbyte)
	nb := make([]byte, k)
	for i := 0; i < k; i++ {
		nb[i] = gbyte[x-i-1]
	}
	return nb
}


//整形转换成字节2位
func IntToBytes2(n int16) []byte {
	//m := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	//使用大端还是小端应根据当前机器进行判断，或者发送方和使用放约定好
	binary.Write(bytesBuffer, binary.LittleEndian, n)

	gbyte := bytesBuffer.Bytes()
	//c++ 高低位转换
	k := 2
	x := len(gbyte)
	nb := make([]byte, k)
	for i := 0; i < k; i++ {
		nb[i] = gbyte[x-i-1]
	}
	return nb
}



