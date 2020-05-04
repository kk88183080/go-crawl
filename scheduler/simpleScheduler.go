package scheduler

import (
	"../engine"
)

type SimpleScheduler struct {
	WorkChan chan engine.Request
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	go func() {
		s.WorkChan <- request
	}()
}

func (s *SimpleScheduler) ConfigWorkChan(r chan engine.Request) {
	s.WorkChan = r
}
