package parser

import (
	"crawler/model"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseProfile(contents, "没那种命")
	fmt.Printf("result: %v", result)
	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1"+"element; but was %v", result.Items)
	}

	profile := result.Items[0].(model.Profile)

	expected := model.Profile{
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
	}

	if profile != expected {
		t.Errorf("expected %v: but was %v ", expected, profile)
	}
}
