package service

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"showcase-go/agent/api"
)

type NodeService interface {
	Register(ctx context.Context, node *api.Node, options metav1.CreateOptions)
	GetMap(ctx context.Context, options metav1.GetOptions) map[string]*api.Node
}

type nodeService struct {
	M map[string]*api.Node
}

func NewNodeService() *nodeService {
	return &nodeService{
		M: make(map[string]*api.Node),
	}
}

func (n *nodeService) Register(ctx context.Context, node *api.Node, options metav1.CreateOptions) {
	name := node.Name
	n.M[name] = node
}

func (n *nodeService) GetMap(ctx context.Context, options metav1.GetOptions) map[string]*api.Node {
	return n.M
}
