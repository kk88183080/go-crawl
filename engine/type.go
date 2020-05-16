/**
数据类型进行定义
*/
package engine

type ParseResult struct {
	Requests []Request
	Items    []Item // 可以存放任何对象
}

type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}

type Parser interface {
	Parse(content []byte, url string) ParseResult // 解析方法
	Serialize() (name string, args interface{})   // 序列化方法
}

type Request struct {
	Url       string
	ParseFunc Parser
}

type NilParse struct {
}

func (f NilParse) Parse(content []byte, url string) ParseResult {
	return ParseResult{}
}

func (f NilParse) Serialize() (name string, args interface{}) {
	return "NilParse", nil
}

type ParseFunc func(content []byte, url string) ParseResult

type FuncParse struct {
	Parser ParseFunc
	Name   string
}

func (f FuncParse) Parse(content []byte, url string) ParseResult {
	return f.Parser(content, url)
}

func (f FuncParse) Serialize() (name string, args interface{}) {
	return f.Name, nil
}

func NewFuncParse(p ParseFunc, name string) *FuncParse {
	return &FuncParse{
		Parser: p,
		Name:   name,
	}
}
