package main

import (
	"../disPersist"
	"../rpcSupport"
	"gopkg.in/olivere/elastic.v5"
)

func main() {
	serveRpc(":1234")
}

func serveRpc(host string) error {
	client, e := elastic.NewClient(elastic.SetSniff(false))

	if e != nil {
		return e
	}

	return rpcSupport.ServerRpc(host, &disPersist.ItemService{
		Client: client,
	})
}
