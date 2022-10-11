package utils

import v1 "k8s.io/api/core/v1"

func GetTotalRequestResource(spec v1.PodSpec, name v1.ResourceName) int64 {
	initContainerRequest := GetRequestResource(spec.InitContainers, name)
	mainContainerRequest := GetRequestResource(spec.Containers, name)
	return initContainerRequest + mainContainerRequest
}

func GetRequestResource(containers []v1.Container, name v1.ResourceName) int64 {
	var total int64 = 0
	for _, c := range containers {
		requestQuantity := c.Resources.Requests[name]
		r, ok := requestQuantity.AsInt64()
		if ok {
			total += r
		}
	}
	return total
}
