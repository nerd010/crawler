package main

import (
	"crawler/crawler-distributed/config"
	"crawler/crawler-distributed/persist/client"
	"crawler/engine"
	"crawler/scheduler"
	"crawler/zhenai/parser"
	"fmt"
)

func main() {
	itemChan, err := client.ItemServer(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, "ParseCityList"),
	})

	//e.Run(engine.Request{
	//	Url:       "http://www.zhenai.com/zhenghun/beijing",
	//	ParseFunc: parser.ParseCity,
	//})
}
