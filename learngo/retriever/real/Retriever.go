package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	UserAgent string
	TimeOut time.Duration
}

func (r *Retriever) Get(url string) string {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	result, err := httputil.DumpResponse(response, true)
	response.Body.Close()
	if err != nil {
		panic(err)
	}
	return string(result)
}

