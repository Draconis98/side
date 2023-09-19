package service

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/pointer"
)

func GetPVC(pvcName, namespace string) (*corev1.PersistentVolumeClaim, error) {
	clientset := GetKubeClient()
	pvcClient, err := clientset.CoreV1().PersistentVolumeClaims(namespace).Get(context.TODO(), pvcName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return pvcClient, nil
}

func CreatePVC(pvc *corev1.PersistentVolumeClaim, namespace string) (*corev1.PersistentVolumeClaim, error) {
	clientset := GetKubeClient()
	pvcClient, err := clientset.CoreV1().PersistentVolumeClaims(namespace).Create(context.TODO(), pvc, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}

	return pvcClient, nil
}

func PVC(namespace string) *corev1.PersistentVolumeClaim {
	return &corev1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "serve-eda-pvc",
			Namespace: namespace,
			Labels: map[string]string{
				"app": namespace,
			},
		},
		Spec: corev1.PersistentVolumeClaimSpec{
			AccessModes: []corev1.PersistentVolumeAccessMode{
				"ReadOnlyMany",
			},
			Resources: corev1.ResourceRequirements{
				Requests: corev1.ResourceList{
					"storage": resource.MustParse("1Gi"),
				},
			},
			StorageClassName: pointer.String(namespace),
		},
	}
}
