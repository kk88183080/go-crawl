package main

import (
	"../api"
	"log"
	"net"
	"net/rpc/jsonrpc"
)

func main() {

	conn, e := net.Dial("tcp", ":1234")

	if e != nil {
		panic(e)
	}

	client := jsonrpc.NewClient(conn)

	var result float64

	e = client.Call("DemoService.Div", api.Args{A: 10, B: 0}, &result)

	if e != nil {
		log.Println(e)
	} else {
		log.Println("result:%v", result)

	}

}
