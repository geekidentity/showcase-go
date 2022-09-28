package node

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

func (n *NodeController) PrintNode(responseWriter http.ResponseWriter, request *http.Request) error {
	nodeMap := n.service.GetMap(context.Background(), metav1.GetOptions{})
	for k, v := range nodeMap {
		fmt.Println(k, v.Cuda.(map[string]interface{})["cuda"])
	}

	return nil
}
