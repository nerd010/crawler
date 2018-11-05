package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(
	`<div class="m-btn purple" data-v-ff544c08="">[\d]+岁</div>`)
var heightRe = regexp.MustCompile(
	`<div class="m-btn purple" data-v-ff544c08="">[\d]+cm</div>`)
var incomeRe = regexp.MustCompile(
	`<div class="m-btn purple" data-v-ff544c08="">月收入:([^<]+)</div>`)

//var weightRe = regexp.MustCompile(
//	`<td><span class="label">体重: </span><span field="">([\d]+)KG</span></td>`)
//var genderRe = regexp.MustCompile(
//	`<td width="180"><span class="grayL">性别：</span>([^<]+)</td>`)
var xinzuoRe = regexp.MustCompile(
	`<div class="m-btn purple" data-v-ff544c08="">([^<]+)</div>`)
var marrageRe = regexp.MustCompile(
	`<div class="m-btn purple" data-v-ff544c08="">([^<]+)</div>`)
var educationRe = regexp.MustCompile(
	`<div class="m-btn purple" data-v-ff544c08="">([^<]+)</div>`)
var occupationRe = regexp.MustCompile(
	`<div class="m-btn purple" data-v-ff544c08="">([^<]+)</div>`)
var hokouRe = regexp.MustCompile(
	`<div class="m-btn purple" data-v-ff544c08="">([^<]+)</div>`)
var houseRe = regexp.MustCompile(
	`<div class="m-btn pink" data-v-ff544c08="">([^<]+)</div>`)
var carRe = regexp.MustCompile(
	`<div class="m-btn pink" data-v-ff544c08="">([^<]+)</div>`)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	// 年龄
	if age, err := strconv.Atoi(extractString(contents, ageRe)); err != nil {
		profile.Age = age
	}
	// 身高
	if height, err := strconv.Atoi(extractString(contents, heightRe)); err != nil {
		profile.Height = height
	}
	// 体重
	//if weight, err := strconv.Atoi(extractString(contents, weightRe)); err != nil {
	//	profile.Weight = weight
	//}
	// 收入
	profile.Income = extractString(contents, incomeRe)
	// 性别
	//profile.Gender = extractString(contents, genderRe)
	// 星座
	profile.Xinzuo = extractString(contents, xinzuoRe)
	// 婚况
	profile.Marriage = extractString(contents, marrageRe)
	// 学历
	profile.Education = extractString(contents, educationRe)
	// 职业
	profile.Occupation = extractString(contents, occupationRe)
	// 籍贯
	profile.Hokou = extractString(contents, hokouRe)
	// 住房条件
	profile.House = extractString(contents, houseRe)
	// 是否购车
	profile.Car = extractString(contents, carRe)

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
