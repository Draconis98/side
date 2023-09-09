package service

import (
	"fmt"
	"log"
	"strings"
	"sync"
)

func Delete(studentName, containerName string) { // studentName is the namespace containerName
	log.Println("Deleting k8s resources...")
	defer fmt.Println("K8s resources deleted")

	containerName = strings.ToLower(containerName)

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		defer log.Printf("Deployment %s \033[31mdeleted\033[0m\n", containerName)

		if err := DeleteDeployment(containerName, studentName); !err {
			log.Printf("Deployment %s \033[33mnot found\033[0m", containerName)
		}
	}() // Delete Deployment

	go func() {
		defer wg.Done()
		defer log.Printf("Service %s \033[31mdeleted\033[0m\n", containerName)

		if err := DeleteService(containerName, studentName); err != nil {
			log.Panicln(err.Error())
		}
	}() // Delete Service

	go func() {
		defer wg.Done()
		defer log.Printf("Ingress %s \033[31mdeleted\033[0m\n", containerName)

		if err := DeleteIngress(containerName, studentName); err != nil {
			log.Panicln(err.Error())
		}
	}() // Delete Ingress

	wg.Wait()
}
