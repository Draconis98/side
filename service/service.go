package service

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func GetService(serviceName, namespace string) (*corev1.Service, error) {
	clientset := GetKubeClient()
	serviceClient, err := clientset.CoreV1().Services(namespace).Get(context.TODO(), serviceName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return serviceClient, nil
}

func CreateService(service *corev1.Service, namespace string) (*corev1.Service, error) {
	clientset := GetKubeClient()
	// Get Service
	serviceClient, err := GetService(service.Name, namespace)
	if err != nil {
		// Create Service
		serviceClient, err = clientset.CoreV1().Services(namespace).Create(context.TODO(), service, metav1.CreateOptions{})
		if err != nil {
			return nil, err
		}
	}

	return serviceClient, nil
}

func DeleteService(serviceName, namespace string) error {
	clientset := GetKubeClient()
	// Get Service
	_, err := GetService(serviceName, namespace)
	if err != nil { // If not exist, return
		return err
	}

	// Delete Service
	if err = clientset.CoreV1().Services(namespace).Delete(context.TODO(), serviceName, metav1.DeleteOptions{}); err != nil {
		return err
	}

	return nil
}

func Service(dpsrName string) *corev1.Service {
	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: dpsrName,
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{
				"app": dpsrName,
			},
			Ports: []corev1.ServicePort{
				{
					Protocol:   corev1.ProtocolTCP,
					Port:       80,
					TargetPort: intstr.FromInt(3000),
				},
			},
		},
	}
}
