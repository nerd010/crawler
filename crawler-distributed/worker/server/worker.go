package main

import (
	"crawler/crawler-distributed/rpcsupport"
	"crawler/crawler-distributed/worker"
	"flag"
	"fmt"
	"log"
)

var port = flag.Int("port", 0, "the port for me to listen on")
func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(rpcsupport.ServerRpc(
		fmt.Sprintf(":%d", *port),
		worker.CrawlService{}))
}
