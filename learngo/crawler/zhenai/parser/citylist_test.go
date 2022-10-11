package parser

import (
	"showcase-go/learngo/crawler/fetcher"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := fetcher.Fetch("https://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	ParseCityList(contents)
}
