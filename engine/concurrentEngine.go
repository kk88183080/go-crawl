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
	log.Println("fetch url:%s", request.Url)

	bodyResult, e := fetch.Fetch(request.Url)

	if e != nil {
		log.Println("fetch error:%s", request.Url)
		return ParseResult{}, e
	}

	return request.ParseFunc(bodyResult), nil
}
