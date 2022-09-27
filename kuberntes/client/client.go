package main

import (
	"context"
	"flag"
	"fmt"
	"k8s.io/api/events/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/klog"
	"path/filepath"
	expire_map "showcase-go/utils"
	"strings"
	"time"
)

func main() {
	event()
	time.Sleep(time.Second * 60 * 60)

}

func event() {

	var kubeconfig *string

	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		return
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	pods, err := clientSet.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

	informerFactory := informers.NewSharedInformerFactory(clientSet, time.Second*30)

	// 创建Informer（相当于注册到工厂中去，这样下面启动的时候就会失去List && Watch对应的资源）
	informer := informerFactory.Events().V1beta1().Events().Informer()
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: func(obj interface{}) {
		event := obj.(*v1beta1.Event)
		processEvent(clientSet, event)
	},
	})

	stopper := make(chan struct{})
	//defer close(stopper) //启动informer, Lister && Watch
	informerFactory.Start(stopper)
	// 等待所有启动的Informer的缓存被同步
	informerFactory.WaitForCacheSync(stopper)
}

// ErrorNode，不能调度的node
type errorNode struct {
	nodeName    string
	createTime  time.Time
	expiredTime int64
}

var expiredMap = expire_map.NewExpiredMap()

func isErrorNode(taskId, nodeName string) bool {
	nodeSet, found := expiredMap.Get(taskId)
	if !found {
		return false
	}
	_, f := nodeSet.(*expire_map.ExpiredMap).Get(nodeName)
	if !f {
		return false
	}
	return true
}

func processEvent(clientSet *kubernetes.Clientset, event *v1beta1.Event) {
	createTimestamp := event.CreationTimestamp
	if time.Now().Unix()-createTimestamp.Unix() > 20 {
		return
	}
	reason := event.Reason
	fmt.Println(reason)
	if strings.TrimSpace(reason) == "UnexpectedAdmissionError" {
		namespace := event.Namespace
		podName := event.Name
		podName = podName[0:strings.LastIndex(podName, ".")]
		nodeName := event.DeprecatedSource.Host
		fmt.Printf("filter plugin %v eventProcessor for podName %v on nodeName %v",
			"Name", podName, nodeName)
		klog.V(3).Infof("filter plugin %v eventProcessor for podName %v on nodeName %v",
			"Name", podName, nodeName)
		pod, err := clientSet.CoreV1().Pods(namespace).Get(context.TODO(), podName, metav1.GetOptions{})
		if err != nil {
			return
		}
		labels := pod.Labels
		if labels == nil {
			fmt.Println("no label", event.ObjectMeta.Labels)
			return
		}
		taskId := labels["taskid"]
		if taskId == "" {
			fmt.Println("no taskId")
			return
		}
		node := errorNode{
			nodeName:    nodeName,
			createTime:  time.Now(),
			expiredTime: time.Now().Unix() + 30,
		}

		nodeSet, _ := expiredMap.GetOrDefault(taskId, expire_map.NewExpiredMap())
		set := nodeSet.(*expire_map.ExpiredMap)
		set.Put(nodeName, node, 10)
		expiredMap.Put(taskId, set, 10)
		print(expiredMap)
		fmt.Println(isErrorNode(taskId, nodeName))
	}
}

func print(m *expire_map.ExpiredMap) {
	fmt.Print("total task ", m.Length())
	m.Foreach(func(key, val interface{}) {
		fmt.Print("taskId ", key, " node ")
		val.(*expire_map.ExpiredMap).Foreach(func(key, val interface{}) {
			fmt.Print(key, " ")
		})
	})
	fmt.Println()

}
