package service

import (
	"flag"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var kubeClient *kubernetes.Clientset

func InitKubeClient() {
	var kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	if home := homeDir(); home != "" {
		*kubeconfig = filepath.Join(home, ".kube", "config")
	}

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	kubeClient = clientset
}

func GetKubeClient() *kubernetes.Clientset {
	return kubeClient
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
