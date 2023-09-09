package service

import (
	"context"
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/utils/pointer"
	"strconv"
	"strings"
)

func GetDeployment(clientset *kubernetes.Clientset, deploymentName, namespace string) (*appsv1.Deployment, error) {
	deploymentClient, err := clientset.AppsV1().Deployments(namespace).Get(context.TODO(), deploymentName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return deploymentClient, nil
}

func CreateDeployment(clientset *kubernetes.Clientset, deployment *appsv1.Deployment, namespace string) (*appsv1.Deployment, error) {
	//// Get Deployment
	//deploymentClient, err := GetDeployment(clientset, deployment.Name, namespace)
	//if err != nil { // If not exist, create it
	// Create Deployment
	deploymentClient, err := clientset.AppsV1().Deployments(namespace).Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	//}

	return deploymentClient, nil
}

func RestoreDeployment(clientset *kubernetes.Clientset, deploymentName string, cpu, memory int) error {
	defer fmt.Printf("Deployment %s \033[32mrestored\033[0m\n", deploymentName)
	parts := strings.Split(deploymentName, "-")
	image := parts[0]
	studentName := parts[2]

	deployment := Deployment(deploymentName, image, strconv.Itoa(cpu), strconv.Itoa(memory)+"Gi", "restore")
	_, err := CreateDeployment(clientset, deployment, studentName)
	if err != nil {
		panic(err.Error())
	}

	return nil
}

func DeleteDeployment(clientset *kubernetes.Clientset, deploymentName, namespace string) bool {
	// Get Deployment
	_, err := GetDeployment(clientset, deploymentName, namespace)
	if err != nil { // If not exist, return
		return false
	}

	// Delete Deployment
	if err = clientset.AppsV1().Deployments(namespace).Delete(context.TODO(), deploymentName, metav1.DeleteOptions{}); err != nil {
		return false
	}

	return true
}

func Deployment(containerName, image, cpu, memory, flag string) *appsv1.Deployment {
	var img string

	if image == "VScode" {
		if flag == "create" {
			img = "gitlab.agileserve.org.cn:15050/zhangsi/sidehub:vscode"
		} else if flag == "restore" {
			img = "gitlab.agileserve.org.cn:15050/zhangsi/sidehub:" + containerName

		}
	}

	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{Name: containerName, Labels: map[string]string{
			"app": containerName,
		}},
		Spec: appsv1.DeploymentSpec{
			Replicas: pointer.Int32(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": containerName,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name: containerName,
					Labels: map[string]string{
						"app": containerName,
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  containerName,
							Image: img,
							Ports: []corev1.ContainerPort{
								{
									ContainerPort: 3000,
								},
							},
							ImagePullPolicy: corev1.PullPolicy("IfNotPresent"),
							Resources: corev1.ResourceRequirements{
								Limits: corev1.ResourceList{
									corev1.ResourceCPU:    resource.MustParse(cpu),
									corev1.ResourceMemory: resource.MustParse(memory),
								},
								Requests: corev1.ResourceList{
									corev1.ResourceCPU:    resource.MustParse(cpu),
									corev1.ResourceMemory: resource.MustParse(memory),
								},
							},
						},
					},
					ImagePullSecrets: []corev1.LocalObjectReference{
						{
							Name: "image-repo",
						},
					},
				},
			},
		},
	}
}

func CheckDeployment(clientset *kubernetes.Clientset, namespace string) (bool, error) {
	dps, err := clientset.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return false, err
	}

	if len(dps.Items) > 0 {
		return true, nil
	}

	return false, nil
}

func ExpandRequirement(clientset *kubernetes.Clientset, deploymentName, namespace string, oldCPU, oldMem, newCPU, newMem int) bool {
	// Get Deployment
	deployment, err := GetDeployment(clientset, deploymentName, namespace)
	if err != nil {
		fmt.Println("Error getting deployment:", err)
		return false
	}

	node := SearchNode(clientset, deploymentName)

	if node == "" {
		fmt.Println("Error searching node for ", deploymentName)
		return false
	}

	flag, pipelineID := TriggerPipeline(node, deploymentName)
	if !flag {
		fmt.Println("Error triggering pipeline for ", deploymentName)
		return false
	}

	if flag = CheckPipelineStatus(pipelineID); !flag {
		fmt.Println("Error checking pipeline status for ", pipelineID, deploymentName)
		return false
	}

	// Expand CPU
	if oldCPU < newCPU {
		deployment.Spec.Template.Spec.Containers[0].Resources.Limits["cpu"] = resource.MustParse(strconv.Itoa(newCPU))
		deployment.Spec.Template.Spec.Containers[0].Resources.Requests["cpu"] = resource.MustParse(strconv.Itoa(newCPU))
	}

	// Expand Memory
	if oldMem < newMem {
		deployment.Spec.Template.Spec.Containers[0].Resources.Limits["memory"] = resource.MustParse(strconv.Itoa(newMem) + "Gi")
		deployment.Spec.Template.Spec.Containers[0].Resources.Requests["memory"] = resource.MustParse(strconv.Itoa(newMem) + "Gi")
	}

	deployment.Spec.Template.Spec.Containers[0].Image = "gitlab.agileserve.org.cn:15050/zhangsi/sidehub:" + deploymentName

	// Update Deployment
	if _, err = clientset.AppsV1().Deployments(namespace).Update(context.TODO(), deployment, metav1.UpdateOptions{}); err != nil {
		return false
	}

	return true
}

func SearchNode(clientset *kubernetes.Clientset, deploymentName string) string {
	namespace := strings.Split(deploymentName, "-")[2]

	pods, err := clientset.CoreV1().Pods(namespace).List(context.Background(), metav1.ListOptions{
		LabelSelector: "app=" + deploymentName,
	})
	if err != nil {
		fmt.Printf("Error searching node for %s: %s\n", deploymentName, err.Error())
		return ""
	}

	for _, pod := range pods.Items {
		if pod.Status.Phase == "Running" {
			fmt.Println("Found node for", deploymentName, ":", pod.Spec.NodeName)
			return pod.Spec.NodeName
		}
	}

	return ""
}
