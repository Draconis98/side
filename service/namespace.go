package service

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetNamespace(clientset *kubernetes.Clientset, namespace string) (*corev1.Namespace, error) {
	namespaceClient, err := clientset.CoreV1().Namespaces().Get(context.TODO(), namespace, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return namespaceClient, nil
}

func CreateNamespace(clientset *kubernetes.Clientset, namespace *corev1.Namespace) (*corev1.Namespace, error) {
	// Get Namespace
	namespaceClient, err := GetNamespace(clientset, namespace.Name)
	if err != nil { // If not exist, create it
		// Create Namespace
		namespaceClient, err = clientset.CoreV1().Namespaces().Create(context.TODO(), &corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: namespace.Name,
			},
		}, metav1.CreateOptions{})
		if err != nil {
			return nil, err
		}
	}

	return namespaceClient, nil
}

func DeleteNamespace(clientset *kubernetes.Clientset, namespace string) error {
	// Get Namespace
	_, err := GetNamespace(clientset, namespace)
	if err != nil { // If not exist, return
		return err
	}

	// Delete Namespace
	if err = clientset.CoreV1().Namespaces().Delete(context.TODO(), namespace, metav1.DeleteOptions{}); err != nil {
		return err
	}

	return nil
}

func Namespace(namespace string) *corev1.Namespace {
	return &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: namespace,
		},
	}
}
