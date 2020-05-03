package parse

import (
	"../engine"
	"regexp"
)

const regexp_book_Str = `<a href="([^"]+)" title="([^"]+)"`

func Parsebook(content []byte) engine.ParseResult {

	compile := regexp.MustCompile(regexp_book_Str)
	submatch := compile.FindAllSubmatch(content, -1)

	result := engine.ParseResult{}
	for _, e := range submatch {
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(e[1]),
			ParseFunc: engine.NilParse,
		})

		result.Items = append(result.Items, string(e[2]))
	}
	return result
}
