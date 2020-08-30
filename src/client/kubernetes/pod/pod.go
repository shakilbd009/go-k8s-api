package pod

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Kpod struct {
	PodNamespace string `json:"pod_namespace"`
	PodName      string `json:"pod_name"`
	Status       string `json:"status"`
	Restarts     int32  `json:"restarts"`
	PodIP        string `json:"pod_ip"`
}

func GetPod(ctx context.Context, client *kubernetes.Clientset) ([]Kpod, error) {

	pods, err := client.CoreV1().Pods(v1.NamespaceAll).List(ctx, v1.ListOptions{})
	if err != nil {
		return nil, err
	}
	result := make([]Kpod, 0, len(pods.Items))
	for _, pod := range pods.Items {
		result = append(result, Kpod{
			PodNamespace: pod.Namespace,
			PodName:      pod.Name,
			Status:       string(pod.Status.Phase),
			Restarts:     pod.Status.ContainerStatuses[0].RestartCount,
			PodIP:        pod.Status.PodIP,
		})
	}
	return result, nil
}
