package client

import (
	"crawler/crawler-distributed/config"
	"crawler/crawler-distributed/rpcsupport"
	"crawler/engine"
	"log"
)

func ItemServer(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item "+"#%d: %v", itemCount, item)
			itemCount++

			result := ""
			err = client.Call(config.ItemSaverRPC, item, &result)
			if err != nil {
				log.Printf("Item Saver: error "+"saving item %v: %v", item, err)
			}
		}
	}()
	return out, nil
}

