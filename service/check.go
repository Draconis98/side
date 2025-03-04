package service

import (
	"log"
	"strings"
	"sync"
)

//"time"
//"log"

// func CheckSuccess(clientset *kubernetes.Clientset, studentName, containerName string) bool {
// 	containerName = strings.ToLower(containerName)

// 	flag := true

// 	var wg sync.WaitGroup
// 	wg.Add(3)

// 	go func() {
// 		defer wg.Done()

// 		_, err := GetDeployment(clientset, containerName, studentName)
// 		if err != nil {
// 			flag = false
// 		}
// 	}()

// 	go func() {
// 		defer wg.Done()

// 		_, err := GetService(clientset, containerName, studentName)
// 		if err != nil {
// 			flag = false
// 		}
// 	}()

// 	go func() {
// 		defer wg.Done()

// 		_, err := GetIngress(clientset, containerName, studentName)
// 		if err != nil {
// 			flag = false
// 		}
// 	}()

// 	wg.Wait()

// 	return flag
// }

func CheckEndLoading(studentName, containerName string) bool {
	containerName = strings.ToLower(containerName)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		flag := false
		for !flag {
			dpt, err := GetDeployment(containerName, studentName)
			if err != nil {
				log.Printf("\033[31mError getting Deployment: %v\033[0m\n\n", err)
			} else {
				if dpt.Status.AvailableReplicas > 0 {
					log.Println("\033[32mDeployment " + containerName + " is ready!\033[0m")
					flag = true
				}
			}
		}
	}()

	go func() {
		defer wg.Done()
		flag := false
		for !flag {
			igs, err := GetIngress(containerName, studentName)
			if err != nil {
				log.Printf("\033[31mError getting Ingress: %v\033[0m\n\n", err)
			} else {
				if igs.Status.LoadBalancer.Ingress != nil {
					log.Println("\033[32mIngress " + containerName + " is ready!\033[0m")
					flag = true
				}
			}
		}
	}()

	wg.Wait()

	return true
}

// func CheckStatus(clientset *kubernetes.Clientset, studentName, image, timestamp string) bool {
// 	name := strings.ToLower(image) + "-" + timestamp + "-" + studentName

// 	dpt, err := GetDeployment(clientset, name, studentName)
// 	if err != nil {
// 		log.Printf("\033[31mError getting Deployment: %v, maybe not exist\033[0m\n\n", err)
// 		return false
// 	}

// 	if dpt.Status.AvailableReplicas > 0 {
// 		return true
// 	}

// 	return false
// }
