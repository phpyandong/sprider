package core

type Request struct{
	Url string
	ParserFunc func([]byte) ParseResult
}
type ParseResult struct {
	Request [] Request
	//Items []interface{}
	Items []Item
}
type Item struct {
	Url string
	Type string
	Id string
	Payload interface{}
}
func NilParser([]byte) ParseResult{
	return ParseResult{}
}