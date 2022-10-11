package parser

import (
	"fmt"
	"regexp"
	"showcase-go/learngo/crawler/engine"
)

const cityListReg = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParserResult {
	reg := regexp.MustCompile(cityListReg)
	matches := reg.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}
	for _, m := range matches {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests,
			engine.Request{
				Url:        string(m[1]),
				ParserFunc: engine.NilParser,
			},
		)
		fmt.Printf("City: %s, URL: %s\n", m[2], m[1])
	}
	fmt.Printf("Matches found: %d\n", len(matches))
	return result
}
