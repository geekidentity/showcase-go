package parser

import (
	"regexp"
	"showcase-go/learngo/crawler/engine"
	"showcase-go/learngo/crawler/model"
	"strconv"
)

var AgeReg = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([\d])+岁</div>`)
var MarriageReg = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>[^<]+</div>`)

func ParseProfile(contents []byte) engine.ParserResult {
	profile := model.Profile{}
	age, err := strconv.Atoi(extractString(contents, AgeReg))
	if err != nil {
		profile.Age = age
	}
	profile.Marriage = extractString(contents, MarriageReg)

	result := engine.ParserResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(contents []byte, reg *regexp.Regexp) string {
	match := reg.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
