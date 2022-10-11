package utils

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"testing"
)

const (
	ResourceNvidiaGpu          v1.ResourceName = "nvidia.com/gpu"
	ResourceTencentVCudaCore   v1.ResourceName = "tencent.com/vcuda-core"
	ResourceTencentVCudaMemory v1.ResourceName = "tencent.com/vcuda-memory"
)

func TestGetTotalRequestResource(t *testing.T) {
	pod := v1.Pod{
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Resources: v1.ResourceRequirements{
						Requests: v1.ResourceList{
							v1.ResourceCPU:           resource.MustParse("1000m"),
							v1.ResourceMemory:        resource.MustParse("1Gi"),
							ResourceNvidiaGpu:        resource.MustParse("1"),
							ResourceTencentVCudaCore: resource.MustParse("200"),
						},
					},
				},
			},
		},
	}
	fmt.Println(GetTotalRequestResource(pod.Spec, ResourceNvidiaGpu))
}
