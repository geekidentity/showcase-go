package node

import (
	"context"
	"encoding/json"
	"io/ioutil"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"net/http"
	"showcase-go/agent/api"
)

func (n *NodeController) Register(responseWriter http.ResponseWriter, request *http.Request) error {
	bytes, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return err
	}
	var node api.Node
	err = json.Unmarshal(bytes, &node)
	log.Println(node.Cuda)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	n.service.Register(context.Background(), &node, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	body := string(bytes)
	log.Println(body)
	_, err = responseWriter.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}
