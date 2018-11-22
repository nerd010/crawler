package main

import (
	"crawler/engine"
	"crawler/persist"
	"crawler/scheduler"
	"crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    persist.ItemServer(),
	}

	e.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList})

	//e.Run(engine.Request{
	//	Url:       "http://www.zhenai.com/zhenghun/beijing",
	//	ParseFunc: parser.ParseCity,
	//})
}
