package view

import (
	"crawler/engine"
	"crawler/frontend/model"
	common "crawler/model"
	"os"
	"testing"
)

func TestSearchResultView_Render(t *testing.T) {
	view := CreateSearchResultView("template.html")

	out, err := os.Create("template.test.html")
	page := model.SearchResult{}
	page.Hits = 123
	item := engine.Item{
		Url:  "http://album.zhenai.com/u/1031157293",
		Type: "zhenai",
		Id:   "1031157293",
		Payload: common.Profile{
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
	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}
	err = view.Render(out, page)
	if err != nil {
		panic(err)
	}

}