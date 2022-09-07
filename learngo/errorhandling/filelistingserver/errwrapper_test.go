package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func errPanic(writer http.ResponseWriter, request *http.Request) error {
	panic("hello")
}

type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()
}

func (e testingUserError) Message() string {
	return string(e)
}

func errUserError(writer http.ResponseWriter, request *http.Request) error {
	panic("hello")
}

func TestErrWrapper(t *testing.T) {
	tests := []struct {
		h       appHandler
		code    int
		message string
	}{
		{errPanic, 500, http.StatusText(http.StatusInternalServerError)},
	}

	for _, tt := range tests {
		f := errWrapper(tt.h)
		responseWriter := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "http://baidu.com", nil)
		f(responseWriter, request)
		b, _ := ioutil.ReadAll(responseWriter.Body)
		body := strings.Trim(string(b), "\n")
		if responseWriter.Code != tt.code || body != tt.message {
			t.Errorf("expect (%d, %s); got (%d, %s)", tt.code, tt.message, responseWriter.Code, body)
		}
	}
}
