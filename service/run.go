package service

import (
	"fmt"
	"k8s.io/client-go/kubernetes"
	"strconv"
	"strings"
	"sync"
)

func Run(clientset *kubernetes.Clientset, studentName, image, timestamp string, cpu, mem int) {
	fmt.Println("Creating k8s resources...")
	defer fmt.Println("K8s resources created")

	name := strings.ToLower(image) + "-" + timestamp + "-" + studentName

	namespace := Namespace(studentName)
	_, err := CreateNamespace(clientset, namespace)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Namespace %s \033[32mcreated\033[0m\n", studentName)

	secret := Secret(studentName)
	_, err = CreateSecret(clientset, secret, studentName)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Secret %s \033[32mcreated\033[0m\n", secret.Name)

	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()

		//var img string

		//if image == "VScode" {
		//	img = "gitlab.oms.agileserve.org.cn:15050/zhangsi/sidehub:vscode"
		//}

		deployment := Deployment(name, image, strconv.Itoa(cpu), strconv.Itoa(mem)+"Gi", "create")
		_, err := CreateDeployment(clientset, deployment, studentName)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("Deployment %s \033[32mcreated\033[0m\n", name)
	}() // Create Deployment

	go func() {
		defer wg.Done()

		service := Service(name)
		_, err := CreateService(clientset, service, studentName)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("Service %s \033[32mcreated\033[0m\n", name)
	}() // Create Service

	go func() {
		defer wg.Done()

		ingress := Ingress(name, name+".oms.agileserve.org.cn")
		_, err := CreateIngress(clientset, ingress, studentName)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("Ingress %s \033[32mcreated\033[0m\n", name)
	}() // Create Ingress

	wg.Wait()
}
