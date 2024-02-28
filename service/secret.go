package service

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetSecret(clientset *kubernetes.Clientset, namespace string) (*corev1.Secret, error) {
	secretClient, err := clientset.CoreV1().Secrets(namespace).Get(context.TODO(), "image-repo", metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return secretClient, nil
}

func CreateSecret(clientset *kubernetes.Clientset, secret *corev1.Secret, namespace string) (*corev1.Secret, error) {
	_, err := GetNamespace(clientset, namespace)
	if err != nil { // If not exist, create it
		log.Printf("Namespace %s \033[31mnot exist\n\033[0m", namespace)
		log.Println("Secret \033[31mcreate failed\033[0m")
		return nil, err
	}

	secretClient, err := GetSecret(clientset, namespace)
	if err != nil { // If not exist, create it
		// Create Secret
		secretClient, err = clientset.CoreV1().Secrets(namespace).Create(context.TODO(), secret, metav1.CreateOptions{})
		if err != nil {
			return nil, err
		}
	}

	return secretClient, nil
}

func DeleteSceret(clientset *kubernetes.Clientset, namespace string) error {
	// Get Secret
	_, err := GetSecret(clientset, namespace)
	if err != nil { // If not exist, return
		return err
	}

	// Delete Secret
	if err = clientset.CoreV1().Secrets(namespace).Delete(context.TODO(), "image-repo", metav1.DeleteOptions{}); err != nil {
		return err
	}

	return nil
}

func Secret(namespace string) *corev1.Secret {
  dockerServer := "https://gitlab.agileserve.org.cn:15050/zhangsi/sidehub"
	// dockerUsername := "loonghan@foxmail.com"
	// dockerPassword := "h2022shtech#"
  dockerServer := "10.30.19.15:30916"
  dockerUsername := "admin"
  dockerPassword := "Harbor12345"

	auth := dockerUsername + ":" + dockerPassword
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(auth))

	return &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "image-repo",
			Namespace: namespace,
		},
		Type: corev1.SecretTypeDockerConfigJson,
		Data: map[string][]byte{
			".dockerconfigjson": []byte(fmt.Sprintf(
				`{"auths":{"%s":{"username":"%s","password":"%s","email":"%s","auth":"%s"}}}`,
				dockerServer, dockerUsername, dockerPassword, dockerUsername, encodedAuth)),
		},
	}
}
