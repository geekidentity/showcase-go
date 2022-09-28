package node

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"net/http"
)

func (n *NodeController) PrintNode(responseWriter http.ResponseWriter, request *http.Request) error {
	nodeMap := n.service.GetMap(context.Background(), metav1.GetOptions{})
	for _, v := range nodeMap {
		log.Println(v)
	}

	return nil
}
