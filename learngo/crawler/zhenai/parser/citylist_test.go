package parser

import (
	"fmt"
	"showcase-go/learngo/crawler/fetcher"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := fetcher.Fetch("https://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s\n", contents)
	result := ParseCityList(contents)
	fmt.Printf("result size : %d", len(result.Requests))
}
