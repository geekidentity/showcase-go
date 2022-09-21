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
	"path/filepath"
	"time"
)

func main() {

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
		reason := event.Reason
		if reason == "UnexpectedAdmissionError" {
			fmt.Println(event.Reason, event.Namespace, event.GetName(), event.DeprecatedSource.Host)
		}
	}})

	stopper := make(chan struct{})
	defer close(stopper) //启动informer, Lister && Watch
	informerFactory.Start(stopper)
	// 等待所有启动的Informer的缓存被同步
	informerFactory.WaitForCacheSync(stopper)
}
