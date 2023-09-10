package getnodes

import (
	"context"
	"fmt"

	"github.com/lostinopensrc/k8s-client-go/kubeconfig"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func GET_NODE_NAME() {
	// create kubeconfig object from path
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig.GET_KUBE_CONFIG())
	if err != nil {
		panic(err.Error())
	}

	// create clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for _, node_name := range nodes.Items {
		fmt.Println("Node name is:", node_name.Name)
	}
}
