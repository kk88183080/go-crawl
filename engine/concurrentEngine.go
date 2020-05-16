package engine

import (
	"../fetch"
	"github.com/go-acme/lego/log"
)

/**
调度器
*/
type Scheduler interface {
	Submit(Request)
	Run()
	WorkReady(chan Request)
	WorkChan() chan Request
}

/**
并发版的
*/
type ConcurrentEngine struct {
	Scheduler Scheduler
	Work      int
	ItemChan  chan Item // 保存数据的通道
}

func (engine *ConcurrentEngine) Run(seed ...Request) {

	out := make(chan ParseResult)

	engine.Scheduler.Run()

	for i := 0; i < engine.Work; i++ {
		CreateWork(engine.Scheduler.WorkChan(), out, engine.Scheduler)
	}

	for _, r := range seed {
		engine.Scheduler.Submit(r)
	}

	for {
		parseResult := <-out

		// 把结果数据打入保存数据的通道
		for _, item := range parseResult.Items {
			go func() {
				engine.ItemChan <- item // 本行代码是阻塞的，所以要加协程
			}()
		}
		for _, r := range parseResult.Requests {
			engine.Scheduler.Submit(r)
		}
	}
}

func CreateWork(in chan Request, out chan ParseResult, s Scheduler) {
	go func() {
		for {
			s.WorkReady(in)
			request := <-in

			parseResult, error := work(request)
			if error != nil {
				continue
			}

			out <- parseResult
		}
	}()
}

func work(request Request) (ParseResult, error) {
	//log.Println("engine work url:%s", request.Url)

	bodyResult, e := fetch.Fetch(request.Url)

	if e != nil {
		log.Println("fetch error:%s", request.Url)
		return ParseResult{}, e
	}

	if request.ParseFunc == nil {
		return ParseResult{}, nil
	}

	log.Println(string(bodyResult))

	return request.ParseFunc(bodyResult), nil
}
