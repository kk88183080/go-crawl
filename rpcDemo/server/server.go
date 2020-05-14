package main

import (
	"../api"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	rpc.Register(api.DemoService{})

	listener, e := net.Listen("tcp", ":1234")

	if e != nil {
		log.Println("rpc server start error !")
		panic(e)
	}

	log.Print("bind", listener.Addr())

	for {
		conn, e := listener.Accept()

		if e != nil {
			log.Println(" accept error ! %v", e)
			continue
		}

		go jsonrpc.ServeConn(conn)
	}
}
