package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(
	`<td><span class="label">年龄: </span>([\d]+)岁</td>`)
var heightRe = regexp.MustCompile(
	`<td><span class="label">身高: </span>([\d]+)CM</td>`)
var incomeRe = regexp.MustCompile(
	`<td><span class="label">月收入: </span>([^<]+)</td>`)
var weightRe = regexp.MustCompile(
	`<td><span class="label">体重: </span><span field="">([\d]+)KG</span></td>`)
var genderRe = regexp.MustCompile(
	`<td><span class="label">性别: </span><span field="">([^<]+)</span></td>`)
var xinzuoRe = regexp.MustCompile(
	`<td><span class="label">星座: </span><span field="">([^<]+)</span></td>`)
var marrageRe = regexp.MustCompile(
	`<td><span class="label">婚况: </span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(
	`<td><span class="label">学历: </span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(
	`<td><span class="label">职业: </span><span field="">([^<]+)</span></td>`)
var hokouRe = regexp.MustCompile(
	`<td><span class="label">籍贯: </span>([^<]+)</td>`)
var houseRe = regexp.MustCompile(
	`<td><span class="label">住房条件: </span><span field="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(
	`<td><span class="label">是否购车: </span><span field="">([^<]+)</span></td>`)

func ParseProfile(contents []byte) engine.ParseResult {
	profile := model.Profile{}

	// 年龄
	if age, err := strconv.Atoi(extractString(contents, ageRe)); err != nil {
		profile.Age = age
	}
	// 身高
	if height, err := strconv.Atoi(extractString(contents, heightRe)); err != nil {
		profile.Height = height
	}
	// 体重
	if weight, err := strconv.Atoi(extractString(contents, weightRe)); err != nil {
		profile.Weight = weight
	}
	// 收入
	profile.Income = extractString(contents, incomeRe)
	// 性别
	profile.Gender = extractString(contents, genderRe)
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
