package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
	"strconv"
	"strings"
)

/*
var ageRe = regexp.MustCompile(
	`<tr><td width="180"><span class="grayL">年龄：</span>([\d]+)</td>`)
var heightRe = regexp.MustCompile(
	`<td width="180"><span class="grayL">身   高：</span>([\d]+)</td>`)
var incomeRe = regexp.MustCompile(
	`<td><span class="grayL">月   薪：</span>([^<]+)</td>`)

//var weightRe = regexp.MustCompile(
//	`<td><span class="label">体重: </span><span field="">([\d]+)KG</span></td>`)
var genderRe = regexp.MustCompile(
	`<tr><td width="180"><span class="grayL">性别：</span>([^<]+)</td>`)

//var xinzuoRe = regexp.MustCompile(
//	`<div class="m-btn purple" data-v-ff544c08="">([^<]+)</div>`)

var marrageRe = regexp.MustCompile(
	`<div class="m-btn purple" data-v-ff544c08="">([^<]+)</div>`)

//var marrageRe = regexp.MustCompile(
//	`<td width="180"><span class="grayL">婚况：</span>([^<]+)</td>`)

//var educationRe = regexp.MustCompile(
//	`<div class="m-btn purple" data-v-ff544c08="">([^<]+)</div>`)
//var occupationRe = regexp.MustCompile(
//	`<div class="m-btn purple" data-v-ff544c08="">([^<]+)</div>`)
var hokouRe = regexp.MustCompile(
	`<td><span class="grayL">居住地：</span>([^<]+)</td></tr>`)

//var houseRe = regexp.MustCompile(
//	`<div class="m-btn pink" data-v-ff544c08="">([^<]+)</div>`)
*/

// Todo 选择合适的方法获取 div 中的数据

type usermModel struct {
	name      string
	age       int
	height    int
	gender    string
	hokou     string
	education string
	marriage  string
	income    string
}

var userBrief = regexp.MustCompile(
	`<div class="des f-cl" data-v-07a0138b>([^<]+)</div>`)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name

	m := splitUserInfo(extractString(contents, userBrief))
	// 年龄
	profile.Age = m.age
	//if age, err := strconv.Atoi(extractString(contents, ageRe)); err != nil {
	////	profile.Age = age
	////}
	// 身高
	profile.Height = m.height
	//if height, err := strconv.Atoi(extractString(contents, heightRe)); err != nil {
	//	profile.Height = height
	//}
	// 体重
	//if weight, err := strconv.Atoi(extractString(contents, weightRe)); err != nil {
	//	profile.Weight = weight
	//}
	// 收入
	profile.Income = m.income
	//profile.Income = extractString(contents, incomeRe)
	// 性别
	//profile.Gender = extractString(contents, genderRe)
	// 星座
	//profile.Xinzuo = extractString(contents, xinzuoRe)
	// 婚况
	profile.Marriage = m.marriage
	//profile.Marriage = extractString(contents, marrageRe)
	// 学历
	profile.Education = m.education
	//profile.Education = extractString(contents, educationRe)
	// 职业
	//profile.Occupation = extractString(contents, occupationRe)
	// 籍贯
	profile.Hokou = m.hokou
	// 住房条件
	//profile.House = extractString(contents, houseRe)
	// 是否购车
	//profile.Car = extractString(contents, carRe)

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

func splitUserInfo(s string) usermModel {
	userinfo := usermModel{}
	m := strings.Split(s, "|")

	if len(m) == 6 {

		if age, err := strconv.Atoi(strings.TrimSuffix(strings.Replace(m[1], " ", "", -1), "岁")); err == nil {

			userinfo.age = age
		}

		if height, err := strconv.Atoi(strings.TrimSuffix(strings.Replace(m[4], " ", "", -1), "cm")); err == nil {

			userinfo.height = height
		}

		userinfo.hokou = strings.Replace(m[0], " ", "", -1)
		userinfo.education = strings.Replace(m[2], " ", "", -1)
		userinfo.marriage = strings.Replace(m[3], " ", "", -1)
		userinfo.income = strings.Replace(m[5], " ", "", -1)
	}

	return userinfo
}
