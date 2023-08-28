package pods

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func GET_POD_NAME() {

	// get kubeconfig1 we cant redifine kubeconfig variable between two packages again .

	var kubeconfig1 *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig1 = flag.String("kubeconfig1", filepath.Join(home, ".kube", "config"), "get kubeconfig1 from home dir")

	} else {
		kubeconfig1 = flag.String("kubeconfig1", "", "provide abosolute path")

	}
	flag.Parse()

	// build current context from kubeconfig1
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig1)

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
