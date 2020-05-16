package client

import (
	"../../engine"
	"../rpcSupport"
	"log"
)

func SaveItem(host string) (chan engine.Item, error) {

	client, e := rpcSupport.NewClient(host)
	if e != nil {
		return nil, e
	}

	out := make(chan engine.Item)

	go func() {
		itemCount := 0
		for {
			item := <-out
			itemCount++
			// 保存数据到数据库
			rpcResult := ""
			callError := client.Call("ItemService.Save", item, &rpcResult)

			if callError != nil {
				log.Printf("rpc call error: %v, %v", item, callError)

			}
		}
	}()

	return out, nil
}
