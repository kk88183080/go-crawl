package scheduler

import (
	"../engine"
)

type QueueScheduler struct {
	requestChan chan engine.Request
	workChan    chan chan engine.Request
}

func (s *QueueScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueueScheduler) ConfigWorkChan(chan engine.Request) {

}

func (s *QueueScheduler) WorkReady(w chan engine.Request) {
	s.workChan <- w
}

func (s *QueueScheduler) Run() {
	var requestQ []engine.Request
	var workQ []chan engine.Request

	go func() {
		for {

			var activeRequest engine.Request
			var activeWork chan engine.Request

			if len(requestQ) > 0 && len(workQ) > 0 {
				activeRequest = requestQ[0]
				activeWork = workQ[0]
			}

			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workChan:
				workQ = append(workQ, w)
			case activeWork <- activeRequest:
				workQ = workQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
