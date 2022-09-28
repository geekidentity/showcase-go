package node

import "showcase-go/agent/server/service"

type NodeController struct {
	service service.NodeService
}

func NewNodeController() *NodeController {
	return &NodeController{
		service: service.NewNodeService(),
	}
}
