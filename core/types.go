package core

import pb "sprider/craw/rpcsupport/proto3"

type ParseFunc func(
	content []byte,
) ParseResult
type Request struct{
	Url string
	ParserFunc ParseFunc
}
type ParseResult struct {
	Request [] Request
	//Items []interface{}
	Items []pb.Item
}
type Item struct {
	//proto3.Item
	Url string
	Type string
	Id string
	Payload *pb.Profile
}
func NilParser([]byte) ParseResult{
	return ParseResult{}
}