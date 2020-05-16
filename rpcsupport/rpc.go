package rpcsupport

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func ServerRpc(host string, service interface{}) error {
	rpc.Register(service)

	listener, e := net.Listen("tcp", host)

	if e != nil {
		return e
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

func NewClient(host string) (*rpc.Client, error) {
	conn, e := net.Dial("tcp", ":1234")

	if e != nil {
		return e, nil
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
