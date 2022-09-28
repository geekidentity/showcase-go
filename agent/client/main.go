package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	fmt.Println(response)
	if err != nil {
		fmt.Println(err.Error())
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
	fmt.Println(node)
	return node
}
