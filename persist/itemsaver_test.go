package persist

import (
	"context"
	"crawler/model"
	"testing"

	"encoding/json"

	"crawler/engine"

	"github.com/olivere/elastic"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
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

	// TODO: Try to start up elastic search
	// here using docker go client.
	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	const index = "dating_test"
	// Save expected item
	err = save(client, index, expected)

	if err != nil {
		panic(err)
	}

	// Fetch saved item
	resp, err := client.Get().
		Index(index).
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	t.Logf("%s", resp.Source)

	var actual engine.Item
	json.Unmarshal(
		*resp.Source, &actual)
	if err != nil {
		panic(err)
	}

	actualProfile, _ := model.FromJsonObj(actual.Payload)

	actual.Payload = actualProfile
	// Verify result
	if actual != expected {
		t.Errorf("got %v; expected %v", actual, expected)
	}
}
