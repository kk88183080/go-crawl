package parse

import (
	"../engine"
	"regexp"
)

//<a href="/tag/小说" class="tag">小说</a>
const regexpStr = `<a href="([^"]+)" class="tag">([^"]+)</a>`

/**
使用正则表达式，匹配所要抓取的内容
*/
func ParseContent(content []byte) engine.ParseResult {

	compile := regexp.MustCompile(regexpStr)
	submatch := compile.FindAllSubmatch(content, -1)

	result := engine.ParseResult{}
	for _, m := range submatch {
		//result.Items = append(result.Items, m[2])
		result.Requests = append(result.Requests, engine.Request{
			Url:       "https://book.douban.com/" + string(m[1]),
			ParseFunc: Parsebook,
		})
	}

	return result
}
