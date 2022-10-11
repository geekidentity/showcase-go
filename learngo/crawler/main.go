package main

import (
	"showcase-go/learngo/crawler/engine"
	"showcase-go/learngo/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "https://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
