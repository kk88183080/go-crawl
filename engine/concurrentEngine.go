package engine

import (
	"../fetch"
	"github.com/go-acme/lego/log"
)

/**
调度器
*/
type Scheduler interface {
	submit(Request)
	configWorkChan(chan Request)
}

type SimpleScheduler struct {
	WorkChan chan Request
}

func (s *SimpleScheduler) submit(request Request) {
	s.WorkChan <- request
}

func (s *SimpleScheduler) configWorkChan(r chan Request) {
	s.WorkChan = r
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

	for i := 0; i < engine.Work; i++ {
		CreateWork(in, out)
	}

	for _, r := range seed {
		//work(r)
		engine.Scheduler.submit(r)
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
