package main

import (
	"crawler/crawler-distributed/config"
	itemsaver "crawler/crawler-distributed/persist/client"
	"crawler/crawler-distributed/rpcsupport"
	worker "crawler/crawler-distributed/worker/client"
	"crawler/engine"
	"crawler/scheduler"
	"crawler/zhenai/parser"
	"flag"
	"log"
	"net/rpc"
	"strings"
)


var (
	itemSaverHost = flag.String(
		"itemsaver_host", "", "itemsaver host")
	workerHosts = flag.String(
		"worker_hosts", "", "worker hosts (comma separated)")
)
func main() {
	flag.Parse()
	itemChan, err := itemsaver.ItemServer(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	pool := createClientPool(strings.Split(*workerHosts, ","))

	processor := worker.CreateProcessor(pool)


	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
		RequestProcessor: processor,
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})

	//e.Run(engine.Request{
	//	Url:       "http://www.zhenai.com/zhenghun/beijing",
	//	ParseFunc: parser.ParseCity,
	//})
}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connected to %s", h)
		} else {
			log.Printf("error connecting to %s: %v", h, err)
		}
	}

	out := make(chan *rpc.Client)
	go func() {
		for  {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}