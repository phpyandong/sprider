package core

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