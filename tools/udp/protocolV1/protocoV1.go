package protocolV1

import (
	"encoding/binary"
	"bytes"
)

/**
	* 包头长度
	* @var integer
	*/
const PACKAGE_FIXED_LENGTH = 17

/**
 * udp 包最大长度
 * @var integer
 */
const MAX_UDP_PACKGE_SIZE  = 65507

/**
 * char类型能保存的最大数值
 * @var integer
 */
const MAX_CHAR_VALUE = 255

/**
 *  usigned short 能保存的最大数值
 * @var integer
 */
const MAX_UNSIGNED_SHORT_VALUE = 65535
/**
 *
 * struct statisticPortocol
 * {
 *     unsigned char module_name_len;
 *     unsigned char interface_name_len;
 *     float cost_time;
 *     unsigned char success;
 *     int code;
 *     unsigned short msg_len;
 *     unsigned int time;
 *     char[module_name_len] module_name;
 *     char[interface_name_len] interface_name;
 *     char[msg_len] msg;
 * }
 *
 * @author workerman.net
 */
 /**
  * @author zhangyandong@taipai.tv
  */
type StatisticProtocol struct {
	//模块名长度
	ModuleNameLen		uint8
	//接口名长度
	InterfaceNameLen  	uint8
	//花费时间
	CostTime			float32
	//是否成功
	Success 			uint8
	//状态码
	Code				int32
	//消息体的长度
	MsgLen				int16
	//请求时间戳
	Time				int32
	//模块名
	ModuleName			string
	//接口名
	InterfaceName		string
	//消息体
	Msg					string

}


func (p *StatisticProtocol)  BytesToStruct(data []byte){
	buf := bytes.NewBuffer(data)
	/*
	*按顺序解码二进制数据到对应的数据
	 */
	//$data = unpack("Cmodule_name_len/Cinterface_name_len/fcost_time/Csuccess/Ncode/nmsg_len/Ntime", $bin_data);
	binary.Read(buf, binary.LittleEndian, &(p.ModuleNameLen))
	binary.Read(buf, binary.LittleEndian, &(p.InterfaceNameLen))
	binary.Read(buf, binary.LittleEndian, &(p.CostTime))
	binary.Read(buf, binary.LittleEndian, &(p.Success))
	binary.Read(buf, binary.BigEndian, &(p.Code))
	binary.Read(buf, binary.BigEndian, &(p.MsgLen))
	binary.Read(buf, binary.BigEndian, &(p.Time))
	//根据moduelnamelen 截取模块名
	p.ModuleName = string(data[PACKAGE_FIXED_LENGTH:PACKAGE_FIXED_LENGTH + p.ModuleNameLen])
	//根据interfaceNameLen 截取接口名
	p.InterfaceName = string(data[PACKAGE_FIXED_LENGTH + p.ModuleNameLen:PACKAGE_FIXED_LENGTH + p.ModuleNameLen+ p.InterfaceNameLen])
	//截取剩余部分作为消息体
	p.Msg = string(data[PACKAGE_FIXED_LENGTH + p.ModuleNameLen+ p.InterfaceNameLen:])
}
func (p *StatisticProtocol) ToBytes() ( []byte,error) {
	buf := new(bytes.Buffer)
	buf.Grow(MAX_UDP_PACKGE_SIZE)
	binary.Write(buf, binary.LittleEndian, p.ModuleNameLen)
	binary.Write(buf, binary.LittleEndian, p.InterfaceNameLen)
	binary.Write(buf, binary.LittleEndian, p.CostTime)
	binary.Write(buf, binary.LittleEndian, p.Success)
	binary.Write(buf, binary.BigEndian, p.Code)
	//buf.WriteByte(byte(p.MsgLen))
	binary.Write(buf, binary.BigEndian, p.MsgLen)
	binary.Write(buf, binary.BigEndian, p.Time)
	//buf.WriteByte(byte(p.Time))
	buf.Write([]byte(p.ModuleName+p.InterfaceName+p.Msg))
	return buf.Bytes(), nil
}