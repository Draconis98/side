package service

import (
	"context"

	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/utils/pointer"
)

func GetIngress(clientset *kubernetes.Clientset, ingressName, namespace string) (*networkingv1.Ingress, error) {
	ingressClient, err := clientset.NetworkingV1().Ingresses(namespace).Get(context.TODO(), ingressName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return ingressClient, nil
}

func CreateIngress(ingress *networkingv1.Ingress, namespace string) (*networkingv1.Ingress, error) {
	clientset := GetKubeClient()
	// Get Ingress
	ingressClient, err := GetIngress(clientset, ingress.Name, namespace)
	if err != nil { // If not exist, create it
		// Create Ingress
		ingressClient, err = clientset.NetworkingV1().Ingresses(namespace).Create(context.TODO(), ingress, metav1.CreateOptions{})
		if err != nil {
			return nil, err
		}
	}

	return ingressClient, nil
}

func DeleteIngress(clientset *kubernetes.Clientset, ingressName, namespace string) error {
	// Get Ingress
	_, err := GetIngress(clientset, ingressName, namespace)
	if err != nil { // If not exist, return
		return err
	}

	// Delete Ingress
	if err = clientset.NetworkingV1().Ingresses(namespace).Delete(context.TODO(), ingressName, metav1.DeleteOptions{}); err != nil {
		return err
	}

	return nil
}

func Ingress(igsrName, url string) *networkingv1.Ingress {
	pathType := networkingv1.PathTypePrefix
	return &networkingv1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name: igsrName,
		},
		Spec: networkingv1.IngressSpec{
			IngressClassName: pointer.String("nginx"),
			Rules: []networkingv1.IngressRule{
				{
					Host: url,
					IngressRuleValue: networkingv1.IngressRuleValue{
						HTTP: &networkingv1.HTTPIngressRuleValue{
							Paths: []networkingv1.HTTPIngressPath{
								{
									PathType: &pathType,
									Path:     "/",
									Backend: networkingv1.IngressBackend{
										Service: &networkingv1.IngressServiceBackend{
											Name: igsrName,
											Port: networkingv1.ServiceBackendPort{
												Number: 80,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
