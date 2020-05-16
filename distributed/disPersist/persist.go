package disPersist

import (
	"../../engine"
	"../../persist"
	"gopkg.in/olivere/elastic.v5"
)

type ItemService struct {
	Client *elastic.Client
}

func (s *ItemService) Save(item engine.Item, result *string) error {
	err := persist.SaveEs(s.Client, item)

	if err == nil {
		*result = "ok"
	}

	return err
}
