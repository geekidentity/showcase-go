package main

import (
	"fmt"
	"learngo/retriever/mock"
	real2 "learngo/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string,
		from map[string]string) string
}

const url = "http://www.imooc.com"

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) string {
	return poster.Post(url,
		map[string]string{
			"name":   "ccmouse",
			"course": "golang",
		})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url,
		map[string]string{
			"contents": "another faked imooc.com",
		})
	return s.Get(url)
}



func main() {
	var r Retriever
	r = &mock.Retriever{Contents: "mock html"}
	fmt.Println(download(r))
	inspect(r)
	r = &real2.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut: time.Minute,
	}
	//fmt.Println(download(r))
	inspect(r)
	// type assertion
	if mockRetriever, ok := r.(*real2.Retriever); ok {
		fmt.Println(mockRetriever.UserAgent)
	}
	mockRetriever := mock.Retriever{
		Contents: "this is a fake imooc.com"}

	fmt.Println(session(&mockRetriever))
}

func inspect(r Retriever) {
	fmt.Printf("%T %v \n", r, r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents: ", v.Contents)
	case *real2.Retriever:
		fmt.Println("UserAgent: ", v.UserAgent)
	}
}


