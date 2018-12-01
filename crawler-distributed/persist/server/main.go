package main

import (
	"crawler/crawler-distributed/persist"
	"crawler/crawler-distributed/rpcsupport"
	"github.com/olivere/elastic"
	"log"
)

func main() {
	log.Fatal(serveRpc(":1234", "dating_profile"))
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}
	return rpcsupport.ServerRpc(host,
		&persist.ItemSaveService{
			Client: client,
			Index: index,
		})
}

