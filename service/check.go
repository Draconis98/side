package service

//"time"
//"fmt"

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

/*

func CheckEndLoading(clientset *kubernetes.Clientset, studentName, image, timestamp string) bool {
	name := strings.ToLower(image) + "-" + timestamp + "-" + studentName

    ingressReady := make(chan struct{})
    deploymentReady := make(chan struct{})

    go func() {
        defer close(ingressReady)
        for {
            igs, err := GetIngress(clientset, name, studentName)
            if err != nil {
                panic(err.Error())
            }
            if igs.Status.LoadBalancer.Ingress != nil {
                break
            }
               time.Sleep(time.Second)
            }
    }()

    go func() {
        defer close(deploymentReady)
        for {
            dpt, err := GetDeployment(clientset, name, studentName)
            if err != nil {
                panic(err.Error())
            }
            if dpt.Status.AvailableReplicas > 0 {
                break
            }
            time.Sleep(time.Second) // 轮询间隔
        }
    }()


    // 等待两个通道都被关闭
    <-ingressReady
    <-deploymentReady

    return true
}



func CheckEndLoading(clientset *kubernetes.Clientset, studentName, image, timestamp string, timeout time.Duration) bool {
	name := strings.ToLower(image) + "-" + timestamp + "-" + studentName

	// 创建超时计时器
	timeoutTimer := time.NewTimer(timeout)
	defer timeoutTimer.Stop()

	ingressReady := make(chan struct{})
	deploymentReady := make(chan struct{})

	go func() {
		defer close(ingressReady)
		for {
			select {
			case <-timeoutTimer.C:
				// 超时时，关闭通道并返回 false
				return
			default:
				igs, err := GetIngress(clientset, name, studentName)
				if err != nil {
					// 处理错误情况，这里可以记录日志或采取其他操作
					fmt.Printf("Error getting Ingress: %v\n", err)
					time.Sleep(time.Second) // 等待一段时间后重试
					continue
				}
				if igs.Status.LoadBalancer.Ingress != nil {
					return
				}
				time.Sleep(time.Second)
			}
		}
	}()

	go func() {
		defer close(deploymentReady)
		for {
			select {
			case <-timeoutTimer.C:
				// 超时时，关闭通道并返回 false
				return
			default:
				dpt, err := GetDeployment(clientset, name, studentName)
				if err != nil {
					// 处理错误情况，这里可以记录日志或采取其他操作
					fmt.Printf("Error getting Deployment: %v\n", err)
					time.Sleep(time.Second) // 等待一段时间后重试
					continue
				}
				if dpt.Status.AvailableReplicas > 0 {
					return
				}
				time.Sleep(time.Second)
			}
		}
	}()

	// 使用 select 监听通道事件
	select {
	case <-ingressReady:
		// Ingress 已经就绪
		// 如果这里需要执行额外的操作，请根据需要添加
	case <-deploymentReady:
		// Deployment 已经就绪
		// 如果这里需要执行额外的操作，请根据需要添加
	case <-timeoutTimer.C:
		// 超时
		return false
	}

	// 如果两个资源都就绪，返回 true
	return true
}

*/

// func CheckEndLoading(clientset *kubernetes.Clientset, studentName, containerName string) bool {
// 	containerName = strings.ToLower(containerName)

// 	var wg sync.WaitGroup
// 	wg.Add(2)

// 	go func() {
// 		defer wg.Done()
// 		flag := false
// 		for !flag {
// 			dpt, err := GetDeployment(clientset, containerName, studentName)
// 			if err != nil {
// 				fmt.Printf("\033[31mError getting Deployment: %v\033[0m\n\n", err)
// 			} else {
// 				if dpt.Status.AvailableReplicas > 0 {
// 					fmt.Println("\033[32mDeployment " + containerName + " is ready!\033[0m")
// 					flag = true
// 				}
// 			}
// 		}
// 	}()

// 	go func() {
// 		defer wg.Done()
// 		flag := false
// 		for !flag {
// 			igs, err := GetIngress(clientset, containerName, studentName)
// 			if err != nil {
// 				fmt.Printf("\033[31mError getting Ingress: %v\033[0m\n\n", err)
// 			} else {
// 				if igs.Status.LoadBalancer.Ingress != nil {
// 					fmt.Println("\033[32mIngress " + containerName + " is ready!\033[0m")
// 					flag = true
// 				}
// 			}
// 		}
// 	}()

// 	wg.Wait()

// 	return true
// }

// func CheckStatus(clientset *kubernetes.Clientset, studentName, image, timestamp string) bool {
// 	name := strings.ToLower(image) + "-" + timestamp + "-" + studentName

// 	dpt, err := GetDeployment(clientset, name, studentName)
// 	if err != nil {
// 		fmt.Printf("\033[31mError getting Deployment: %v, maybe not exist\033[0m\n\n", err)
// 		return false
// 	}

// 	if dpt.Status.AvailableReplicas > 0 {
// 		return true
// 	}

// 	return false
// }
