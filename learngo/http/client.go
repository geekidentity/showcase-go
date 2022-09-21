package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func basic() {
	response, err := http.Get("https://www.baidu.com")
	if err != nil {
		return
	}
	defer response.Body.Close()

	dumpResponse, err := httputil.DumpResponse(response, true)
	if err != nil {
		return
	}
	fmt.Println(string(dumpResponse))
}

func get() {
	request, err := http.NewRequest(http.MethodGet, "http://www.imooc.com", nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Mobile Safari/537.36")
	if err != nil {
		return
	}
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect:", req)
			return nil
		},
	}
	response, err := client.Do(request)
	defer response.Body.Close()
	if err != nil {
		return
	}
	dumpResponse, err := httputil.DumpResponse(response, true)
	if err != nil {
		return
	}
	fmt.Println(string(dumpResponse))
}

func main() {
	get()
}
