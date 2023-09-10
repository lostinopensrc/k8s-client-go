package pods

import (
	"context"
	"fmt"

	"github.com/lostinopensrc/k8s-client-go/kubeconfig"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func GET_POD_NAME() {

	// create kubeconfig object from path
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig.GET_KUBE_CONFIG())

	if err != nil {
		panic(err.Error())
	}

	// build clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// get pods in a namespace
	ns := "tigera-operator"
	pods, err := clientset.CoreV1().Pods(ns).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for _, pod := range pods.Items {
		fmt.Println("The pods in namespace tigera-operator are:", pod.Name)
	}

}
