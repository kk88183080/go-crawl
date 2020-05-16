package persist

import (
	"../engine"
	"context"
	"errors"
	"gopkg.in/olivere/elastic.v5"
)

func SaveItem() (chan engine.Item, error) {

	client, e := elastic.NewClient(elastic.SetSniff(false))
	if e != nil {
		return nil, e
	}

	out := make(chan engine.Item)

	go func() {
		itemCount := 0
		for {
			item := <-out
			//log.Printf("save item $%d, %v", itemCount, item)
			itemCount++
			// 保存数据到数据库
			saveEs(client, item)
			//saveMysql(item)
		}
	}()

	return out, nil
}

func saveMysql(item engine.Item) {

}

func saveEs(client *elastic.Client, item engine.Item) error {

	if item.Type == "" {
		return errors.New("type is not null or empty")
	}

	indexService := client.Index().Index("db_name_1").Type(item.Type).BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err := indexService.Do(context.Background())

	if err != nil {
		panic(err)
	}

	return nil
}
