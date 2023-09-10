package kubeconfig

import (
	"fmt"
	"os"
	"path/filepath"
)

func GET_KUBE_CONFIG() *string {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("error getting user home directory %v\n", err)
		os.Exit(1)
	}
	KubeConfigPath := filepath.Join(userHomeDir, ".kube", "config")
	fmt.Printf("Using kubeconfigPath: %s\n", KubeConfigPath)

	return &KubeConfigPath
}
