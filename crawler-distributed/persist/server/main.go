package main

import (
	"crawler/crawler-distributed/config"
	"crawler/crawler-distributed/persist"
	"crawler/crawler-distributed/rpcsupport"
	"fmt"
	"github.com/olivere/elastic"
	"log"
)

func main() {
	log.Fatal(serveRpc(fmt.Sprintf(":%d", config.ItemSaverPort), config.ElasticIndex))
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

