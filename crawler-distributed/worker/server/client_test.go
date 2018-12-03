package main

import (
	"crawler/crawler-distributed/config"
	"crawler/crawler-distributed/rpcsupport"
	"crawler/crawler-distributed/worker"
	"fmt"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServerRpc(host, worker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	req := worker.Request{
		Url: "http://album.zhenai.com/u/1031157293",
		Parser: worker.SerializedParser{
			Name: config.ParseProfile,
			Args: "没那种命",
		},
	}

	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)

	if err != nil {
		t.Error (err)
	} else {
		fmt.Println(result)
	}
}