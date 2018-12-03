package main

import (
	"crawler/crawler-distributed/config"
	"crawler/crawler-distributed/persist"
	"crawler/crawler-distributed/rpcsupport"
	"flag"
	"fmt"
	"github.com/olivere/elastic"
	"log"
)

var port = flag.Int("port", 0, "the port for me to listen on")
func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(serveRpc(fmt.Sprintf(":%d", *port), config.ElasticIndex))
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

