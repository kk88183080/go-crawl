package rpcSupport

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func ServerRpc(host string, servie interface{}) error {
	rpc.Register(servie)

	listener, e := net.Listen("tcp", host)

	if e != nil {
		log.Println("rpc server start error !")
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

	return nil
}

func NewClient(host string) (*rpc.Client, error) {
	conn, e := net.Dial("tcp", host)

	if e != nil {
		return nil, e
	}

	client := jsonrpc.NewClient(conn)

	return client, nil
}
