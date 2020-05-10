package main

import (
	"./engine"
	"./parse/zhenai"
	"./persist"
	"./scheduler"
)

func main() {

	// 爬取tag列表
	//engine.Run(engine.Request{
	//	Url:       "https://book.douban.com/",
	//	ParseFunc: parse.ParseContent,
	//})

	// 爬取book列表
	//engine.Run(engine.Request{
	//	Url:       "https://book.douban.com/tag/%E7%A5%9E%E7%BB%8F%E7%BD%91%E7%BB%9C",
	//	ParseFunc: parse.Parsebook,
	//})

	// 爬取book详情页面
	//engine.Run(engine.Request{
	//	Url:       "https://book.douban.com/subject/30293801/",
	//	ParseFunc: parse.ParseDetailContent,
	//})
	//
	// 爬取book详情页面
	//engine.Run(engine.Request{
	//	Url: "https://book.douban.com/subject/6082808/",
	//	ParseFunc: func(bytes []byte) engine.ParseResult {
	//		return parse.ParseDetailContent(bytes, "百年孤独")
	//	},
	//})

	// 并发版的
	//e := engine.ConcurrentEngine{
	//	Scheduler: &scheduler.SimpleScheduler{},
	//	Work:      100,
	//}
	//
	//e.Run(engine.Request{
	//	Url:       "https://book.douban.com/",
	//	ParseFunc: parse.ParseContent,
	//})

	// 并发版的
	//e := engine.ConcurrentEngine{
	//	Scheduler: &scheduler.SimpleScheduler{},
	//	Work:      100,
	//	ItemChan:  persist.SaveItem(),
	//}
	//
	//e.Run(engine.Request{
	//	Url:       "https://book.douban.com/",
	//	ParseFunc: parse.ParseContent,
	//})

	// 并发版的
	//e := engine.ConcurrentEngine{
	//	Scheduler: &scheduler.SimpleScheduler{},
	//	Work:      100,
	//	ItemChan: persist.SaveItem(),
	//}
	//
	//e.Run(engine.Request{
	//	Url:       "http://www.zhenai.com/zhenghun",
	//	ParseFunc: zhenai.ParseCity,
	//})

	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		Work:      100,
		ItemChan:  persist.SaveItem(),
	}

	e.Run(engine.Request{
		Url: "https://album.zhenai.com/u/1045778053",
		ParseFunc: func(bytes []byte) engine.ParseResult {
			return zhenai.ParsePersonDetail(bytes, "阿里征婚", "女")
		},
	})
	/*e.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun/ali",
		ParseFunc: func(bytes []byte) engine.ParseResult {
			return zhenai.ParsePersonList(bytes, "阿里征婚")
		},
	})*/
}
