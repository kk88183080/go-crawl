package persist

import (
	"context"
	"github.com/go-acme/lego/log"
	"gopkg.in/olivere/elastic.v5"
)

func SaveItem() chan interface{} {
	out := make(chan interface{})

	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("save item $%d, %v", itemCount, item)
			itemCount++
			// 保存数据到数据库
			saveEs(item)
		}

	}()

	return out
}

func saveEs(item interface{}) {
	client, e := elastic.NewClient(
		elastic.SetSniff(false))
	if e != nil {
		panic(e)
	}

	_, err := client.Index().Index("db_name_1").Type("table_name_1").BodyJson(item).Do(context.Background())

	if err != nil {
		panic(err)
	}
}
