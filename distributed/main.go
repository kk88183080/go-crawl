package main

import (
	"../engine"
	"../parse/zhenai"
	"../scheduler"
	"./client"
)

func main() {
	// 如果保存函数没有创建成功，直接退出
	items, err := client.SaveItem(":1234")

	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		Work:      100,
		ItemChan:  items,
	}

	e.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: engine.NewFuncParse(zhenai.ParseCity, "ParseCity"),
	})
}
