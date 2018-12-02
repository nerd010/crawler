package main

import (
	"crawler/crawler-distributed/config"
	"crawler/crawler-distributed/rpcsupport"
	"crawler/engine"
	"crawler/model"
	"fmt"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"

	// start ItemSaverServer
	go serveRpc(fmt.Sprintf(":%d", config.ItemSaverPort), "test1")
	time.Sleep(time.Second)
	// start ItemSaverClient
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	// Call
	item := engine.Item{
		Url:  "http://album.zhenai.com/u/1031157293",
		Type: "zhenai",
		Id:   "1031157293",
		Payload: model.Profile{
			Age:    35,
			Height: 165,
			//Weight:     57,
			Income: "3001-5000元",
			//Xinzuo:    "魔羯座",
			Education: "高中及以下",
			Name:      "没那种命",
			//Gender:     "男",
			Marriage: "离异",
			//Occupation: "销售总监",
			Hokou: "阿坝",
			//House:      "和家人同住",
			//Car: "未买车",
		},
	}

	result := ""
	err = client.Call(config.ItemSaverRPC, item, &result)

	if err != nil || result != "ok" {
		t.Errorf("result: %s, err: %s", result, err)
	}

}