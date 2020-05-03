package engine

import (
	"../fetch"
	"fmt"
	"github.com/go-acme/lego/log"
)

func Run(seed ...Request) {

	var request []Request

	for _, e := range seed {
		request = append(request, e)
	}

	for len(request) > 0 {

		r := request[0]

		request = request[1:]

		log.Println("fetch url:%s", r.Url)
		bodyResult, e := fetch.Fetch(r.Url)

		if e != nil {
			log.Println("fetch error:%s", r.Url)
			panic(e)
		}

		parseResult := r.ParseFunc(bodyResult)
		request = append(request, parseResult.Requests...)

		for _, item := range parseResult.Items {
			fmt.Printf("got item %s\n", item)
		}
	}
}
