package service

import (
	"fmt"
	"k8s.io/client-go/kubernetes"
	"strings"
	"sync"
)

func Delete(clientset *kubernetes.Clientset, studentName, containerName string) { // studentName is the namespace containerName
	fmt.Println("Deleting k8s resources...")
	defer fmt.Println("K8s resources deleted")

	containerName = strings.ToLower(containerName)

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		defer fmt.Printf("Deployment %s \033[31mdeleted\033[0m\n", containerName)

		if err := DeleteDeployment(clientset, containerName, studentName); !err {
			fmt.Printf("Deployment %s \033[33mnot found\033[0m", containerName)
		}
	}() // Delete Deployment

	go func() {
		defer wg.Done()
		defer fmt.Printf("Service %s \033[31mdeleted\033[0m\n", containerName)

		if err := DeleteService(clientset, containerName, studentName); err != nil {
			panic(err.Error())
		}
	}() // Delete Service

	go func() {
		defer wg.Done()
		defer fmt.Printf("Ingress %s \033[31mdeleted\033[0m\n", containerName)

		if err := DeleteIngress(clientset, containerName, studentName); err != nil {
			panic(err.Error())
		}
	}() // Delete Ingress

	wg.Wait()
}
