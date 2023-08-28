package main

import (
	"github.com/lostinopensrc/k8s-client-go/getnodes"
	"github.com/lostinopensrc/k8s-client-go/pods"
)

func main() {

	getnodes.GET_NODE_NAME()
	pods.GET_POD_NAME()
}
