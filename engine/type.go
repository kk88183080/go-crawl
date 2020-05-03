/**
数据类型进行定义
*/
package engine

type ParseResult struct {
	Requests []Request
	Items    []interface{} // 可以存放任何对象
}

type Request struct {
	Url       string
	ParseFunc func([]byte) ParseResult
}

func NilParse([]byte) ParseResult {
	return ParseResult{}
}
