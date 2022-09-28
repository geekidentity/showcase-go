package main

import (
	"log"
	"net/http"
	"os"
	"showcase-go/agent/server/controller/node"
)

func initRouter() {
	nodeController := node.NewNodeController()
	http.HandleFunc("/register", errWrapper(nodeController.Register))
	http.HandleFunc("/printNode", errWrapper(nodeController.PrintNode))

}

type userError interface {
	error
	Message() string
}

type appHandler func(responseWriter http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(responseWriter http.ResponseWriter, request *http.Request) {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(responseWriter,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
				return
			}
		}()
		err := handler(responseWriter, request)
		if err != nil {
			log.Printf("Error handling request: %s", err.Error())

			// user error
			if userErr, ok := err.(userError); ok {
				http.Error(responseWriter, userErr.Message(), http.StatusBadRequest)
				return
			}

			// system error
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden

			default:
				code = http.StatusInternalServerError
			}

			http.Error(responseWriter, http.StatusText(code), code)
		}
	}
}
