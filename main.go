package main

import (
	"./engine"
	"./parse"
)

func main() {

	engine.Run(engine.Request{
		Url:       "https://book.douban.com/",
		ParseFunc: parse.ParseContent,
	})
}
