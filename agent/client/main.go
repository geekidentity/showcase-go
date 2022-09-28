package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"showcase-go/agent/api"
	"strings"
)

var M = make(map[string]interface{})

func main() {
	node := nodeInfo()
	marshal, err := json.Marshal(node)
	if err != nil {
		return
	}
	body := strings.NewReader(string(marshal))
	request, err := http.NewRequest(http.MethodPost, "http://localhost:8888/register", body)
	if err != nil {
		return
	}
	client := http.DefaultClient
	response, err := client.Do(request)
	log.Println(response)
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func nodeInfo() api.Node {
	fileName, _ := os.LookupEnv("filename")
	bytes, _ := ioutil.ReadFile(fileName)
	hostName, _ := os.LookupEnv("Hostname")
	internalIP, _ := os.LookupEnv("InternalIP")
	json.Unmarshal(bytes, &M)
	var node = api.Node{
		Name: hostName,
		Addresses: api.Addresses{
			Hostname:   hostName,
			InternalIP: internalIP,
		},
		Cuda: M["cuda"],
	}
	log.Println(node)
	return node
}
