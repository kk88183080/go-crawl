package api

import (
	"errors"
	"log"
)

type DemoService struct {
}

type Args struct {
	A, B int
}

func (DemoService) Div(para Args, result *float64) error {

	if para.B == 0 {
		return errors.New("B is zero")
	}

	*result = float64(para.A) / float64(para.B)

	log.Println("server div %v", &*result)

	return nil
}
