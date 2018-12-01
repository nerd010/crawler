package persist

import (
	"crawler/engine"
	"crawler/persist"
	"github.com/olivere/elastic"
)

type ItemSaveService struct {
	Client *elastic.Client
 	Index string
}

func (s *ItemSaveService) Save(item engine.Item, result *string) error {
	err := persist.Save(s.Client, s.Index, item)
	if err == nil {
		*result = "ok"
	}
	return err
}