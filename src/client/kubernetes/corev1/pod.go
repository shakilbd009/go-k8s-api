package corev1

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var (
	Pod podIinterface = &pod{}
)

type podIinterface interface {
	GetPods(context.Context, *kubernetes.Clientset, string) (*corev1.PodList, error)
}
type pod struct{}

func (*pod) GetPods(ctx context.Context, client *kubernetes.Clientset, namespace string) (*corev1.PodList, error) {

	pods, err := client.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return pods, nil
}
