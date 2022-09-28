package main

import (
	"net/http"
)

func main() {

	initRouter()

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		return
	}
}
