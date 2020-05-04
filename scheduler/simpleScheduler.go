package scheduler

import (
	"../engine"
)

type SimpleScheduler struct {
	WorkerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	go func() {
		s.WorkerChan <- request
	}()
}

func (s *SimpleScheduler) Run() {
	s.WorkerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) WorkReady(w chan engine.Request) {
	return
}

func (s *SimpleScheduler) WorkChan() chan engine.Request {
	return s.WorkerChan
}
