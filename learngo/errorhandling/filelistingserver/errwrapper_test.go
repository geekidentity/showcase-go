package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
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
	return testingUserError("user error")
}

func errNotFoundError(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrNotExist
}

func errNoPermission(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrPermission
}

func errUnknown(writer http.ResponseWriter, request *http.Request) error {
	return errors.New("unknown error")
}

func noError(writer http.ResponseWriter, request *http.Request) error {
	fmt.Fprintln(writer, "no error")
	return nil
}

var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	{errPanic, 500, http.StatusText(http.StatusInternalServerError)},
	{errUserError, 400, "user error"},
	{errNotFoundError, http.StatusNotFound, http.StatusText(http.StatusNotFound)},
	{errNoPermission, http.StatusForbidden, http.StatusText(http.StatusForbidden)},
	{errUnknown, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)},
	{noError, http.StatusOK, "no error"},
}

func TestErrWrapper(t *testing.T) {

	for _, tt := range tests {
		f := errWrapper(tt.h)
		responseWriter := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "http://baidu.com", nil)
		f(responseWriter, request)
		verifyResponse(responseWriter.Result(), tt.code, tt.message, t)
	}
}

func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		response, _ := http.Get(server.URL)
		verifyResponse(response, tt.code, tt.message, t)
	}
}

func verifyResponse(response *http.Response,
	expectedCode int,
	expectedMessage string,
	t *testing.T,
) {
	b, _ := ioutil.ReadAll(response.Body)
	body := strings.Trim(string(b), "\n")
	if response.StatusCode != expectedCode || body != expectedMessage {
		t.Errorf("expect (%d, %s); got (%d, %s)", expectedCode, expectedMessage, response.StatusCode, body)
	}
}
