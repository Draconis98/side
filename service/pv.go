package service

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreatePV(pv *corev1.PersistentVolume) (*corev1.PersistentVolume, error) {
	clientset := GetKubeClient()
	pvClient, err := clientset.CoreV1().PersistentVolumes().Create(context.TODO(), pv, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}

	return pvClient, nil
}

func PV(namespace string) *corev1.PersistentVolume {
	return &corev1.PersistentVolume{
		ObjectMeta: metav1.ObjectMeta{
			Name: "vivado-" + namespace,
		},
		Spec: corev1.PersistentVolumeSpec{
			AccessModes: []corev1.PersistentVolumeAccessMode{
				"ReadOnlyMany",
			},
			Capacity: corev1.ResourceList{
				"storage": resource.MustParse("2Gi"),
			},
			PersistentVolumeReclaimPolicy: corev1.PersistentVolumeReclaimRetain,
			StorageClassName:              namespace,
			PersistentVolumeSource: corev1.PersistentVolumeSource{
				NFS: &corev1.NFSVolumeSource{
					Server: "10.30.20.100",
					Path:   "/vivado_test",
				},
			},
		},
	}
}
