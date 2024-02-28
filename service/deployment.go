package service

import (
	"context"
	"log"
	"strconv"
	"strings"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/utils/pointer"
)

func GetDeployment(deploymentName, namespace string) (*appsv1.Deployment, error) {
	clientset := GetKubeClient()
	deploymentClient, err := clientset.AppsV1().Deployments(namespace).Get(context.TODO(), deploymentName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return deploymentClient, nil
}

func CreateDeployment(deployment *appsv1.Deployment, namespace string) (*appsv1.Deployment, error) {
	clientset := GetKubeClient()
	// Create Deployment
	deploymentClient, err := clientset.AppsV1().Deployments(namespace).Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	//}

	return deploymentClient, nil
}

func RestoreDeployment(clientset *kubernetes.Clientset, deploymentName string, cpu, memory int) error {
	defer log.Printf("Deployment %s \033[32mrestored\033[0m\n", deploymentName)
	parts := strings.Split(deploymentName, "-")
	image := parts[0]
	studentName := parts[2]

	deployment := Deployment(deploymentName, image, strconv.Itoa(cpu), strconv.Itoa(memory)+"Gi", "restore")
	_, err := CreateDeployment(deployment, studentName)
	if err != nil {
		panic(err.Error())
	}

	return nil
}

func DeleteDeployment(deploymentName, namespace string) bool {
	clientset := GetKubeClient()
	// Get Deployment
	_, err := GetDeployment(deploymentName, namespace)
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

	if image == "vscode" {
		// if flag == "create" {
    img = "10.30.19.15:30916/side/vscode:latest"
		// } else if flag == "restore" {
		// img = "gitlab.agileserve.org.cn:15050/zhangsi/sidehub:" + containerName
		// }
	} else if image == "cod" {
		img = "gitlab.agileserve.org.cn:15050/zhangsi/sidehub:" + strings.Split(containerName, "-")[2]
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
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      "serve-eda-pvc",
									MountPath: "/opt",
									ReadOnly:  true,
								},
							},
						},
					},
					ImagePullSecrets: []corev1.LocalObjectReference{
						{
							Name: "image-repo",
						},
					},
					Volumes: []corev1.Volume{
						{
							Name: "serve-eda-pvc",
							VolumeSource: corev1.VolumeSource{
								PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
									ClaimName: "serve-eda-pvc",
								},
							},
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

func ExpandRequirement(deploymentName, namespace string, oldCPU, oldMem, newCPU, newMem int) bool {
	clientset := GetKubeClient()
	// Get Deployment
	deployment, err := GetDeployment(deploymentName, namespace)
	if err != nil {
		log.Println("Error getting deployment:", err)
		return false
	}

	node := SearchNode(deploymentName)

	if node == "" {
		log.Println("Error searching node for ", deploymentName)
		return false
	}

	flag, pipelineID := TriggerPipeline(node, deploymentName)
	if !flag {
		log.Println("Error triggering pipeline for ", deploymentName)
		return false
	}

	if flag = CheckPipelineStatus(pipelineID); !flag {
		log.Println("Error checking pipeline status for ", pipelineID, deploymentName)
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

  deployment.Spec.Template.Spec.Containers[0].Image = "10.30.19.15:30916/side/" + strings.Split(deploymentName, "-")[0] + deploymentName[strings.Index(deploymentName,"-")+1:]

	// Update Deployment
	if _, err = clientset.AppsV1().Deployments(namespace).Update(context.TODO(), deployment, metav1.UpdateOptions{}); err != nil {
		return false
	}

	return true
}

func SearchNode(deploymentName string) string {
	clientset := GetKubeClient()
	namespace := strings.Split(deploymentName, "-")[2]

	pods, err := clientset.CoreV1().Pods(namespace).List(context.Background(), metav1.ListOptions{
		LabelSelector: "app=" + deploymentName,
	})
	if err != nil {
		log.Printf("Error searching node for %s: %s\n", deploymentName, err.Error())
		return ""
	}

	for _, pod := range pods.Items {
		if pod.Status.Phase == "Running" {
			log.Println("Found node for", deploymentName, ":", pod.Spec.NodeName)
			return pod.Spec.NodeName
		}
	}

	return ""
}
