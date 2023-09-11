package service

import (
	"log"
	"strconv"
	"strings"
	"sync"
)

func Run(studentName, image, timestamp string, cpu, mem int) {
	clientset := GetKubeClient()
	log.Println("Creating k8s resources...")
	defer log.Println("K8s resources created")

	name := strings.ToLower(image) + "-" + timestamp + "-" + studentName

	namespace := Namespace(studentName)
	_, err := CreateNamespace(clientset, namespace)
	if err != nil {
		panic(err.Error())
	}
	log.Printf("Namespace %s \033[32mcreated\033[0m\n", studentName)

	secret := Secret(studentName)
	_, err = CreateSecret(clientset, secret, studentName)
	if err != nil {
		panic(err.Error())
	}
	log.Printf("Secret %s \033[32mcreated\033[0m\n", secret.Name)

	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()

		deployment := Deployment(name, image, strconv.Itoa(cpu), strconv.Itoa(mem)+"Gi", "create")
		_, err := CreateDeployment(deployment, studentName)
		if err != nil {
			panic(err.Error())
		}
		log.Printf("Deployment %s \033[32mcreated\033[0m\n", name)
	}() // Create Deployment

	go func() {
		defer wg.Done()

		service := Service(name)
		_, err := CreateService(service, studentName)
		if err != nil {
			panic(err.Error())
		}
		log.Printf("Service %s \033[32mcreated\033[0m\n", name)
	}() // Create Service

	go func() {
		defer wg.Done()

		ingress := Ingress(name, name+".side.agileserve.org.cn")
		_, err := CreateIngress(ingress, studentName)
		if err != nil {
			panic(err.Error())
		}
		log.Printf("Ingress %s \033[32mcreated\033[0m\n", name)
	}() // Create Ingress

	wg.Wait()
}
