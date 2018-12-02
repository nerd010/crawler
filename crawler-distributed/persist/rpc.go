package persist

import (
	"crawler/engine"
	"crawler/persist"
	"github.com/olivere/elastic"
	"log"
)

type ItemSaveService struct {
	Client *elastic.Client
 	Index string
}

func (s *ItemSaveService) Save(item engine.Item, result *string) error {
	err := persist.Save(s.Client, s.Index, item)
	log.Printf("Item %v saved.", item)
	if err == nil {
		*result = "ok"
	} else {
		log.Printf("Error saving Item %v: %v", item, err)
	}
	return err
}