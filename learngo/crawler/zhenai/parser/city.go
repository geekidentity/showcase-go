package parser

import (
	"fmt"
	"regexp"
	"showcase-go/learngo/crawler/engine"
)

var cityReg = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParserResult {
	reg := regexp.MustCompile(cityReg)
	matches := reg.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}
	for _, m := range matches {
		name := string(m[2])
		result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests,
			engine.Request{
				Url: string(m[1]),
				ParserFunc: func(bytes []byte) engine.ParserResult {
					return ParseProfile(bytes, name)
				},
			},
		)
		fmt.Printf("City: %s, URL: %s\n", m[2], m[1])
	}
	fmt.Printf("Matches found: %d\n", len(matches))
	return result
}
