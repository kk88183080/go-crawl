package zhenai

import (
	"../../engine"
	"regexp"
)

const city_reg = `<a href="([^"]+)" data-v-2cb5b6a2>([^"]+)</a>`

func ParseCity(content []byte) engine.ParseResult {
	compile := regexp.MustCompile(city_reg)

	submatch := compile.FindAllSubmatch(content, -1)
	//log.Println("city name:%v", submatch)

	result := engine.ParseResult{}

	for _, rs := range submatch {
		//log.Println("city name:%s, url:%s", string(rs[1]), string(rs[2]))
		result.Items = append(result.Items, rs[2])
		result.Requests = append(result.Requests, engine.Request{
			Url: string(rs[1]),
			ParseFunc: func(bytes []byte) engine.ParseResult {
				return ParsePersonList(bytes, string(rs[2]))
			},
		})

	}
	return result
}
