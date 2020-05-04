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
	ConfigWorkChan(chan Request)
}

/**
并发版的
*/
type ConcurrentEngine struct {
	Scheduler Scheduler
	Work      int
}

func (engine *ConcurrentEngine) Run(seed ...Request) {

	in := make(chan Request)
	out := make(chan ParseResult)

	engine.Scheduler.ConfigWorkChan(in)

	for i := 0; i < engine.Work; i++ {
		CreateWork(in, out)
	}

	for _, r := range seed {
		engine.Scheduler.Submit(r)
	}

	for {
		parseResult := <-out

		for _, r := range parseResult.Requests {
			engine.Scheduler.Submit(r)
		}
	}
}

func CreateWork(in chan Request, out chan ParseResult) {
	go func() {
		for {
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
	log.Println("fetch url:%s", request.Url)

	bodyResult, e := fetch.Fetch(request.Url)

	if e != nil {
		log.Println("fetch error:%s", request.Url)
		return ParseResult{}, e
	}

	return request.ParseFunc(bodyResult), nil
}
