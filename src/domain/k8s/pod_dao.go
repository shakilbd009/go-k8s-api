package k8s

import (
	"context"

	"github.com/shakilbd009/go-k8s-api/src/client/kubernetes/corev1"
	"github.com/shakilbd009/go-utils-lib/rest_errors"
	"k8s.io/client-go/kubernetes"
)

func (k *Kpod) GetPods(ctx context.Context, client *kubernetes.Clientset) ([]Kpod, rest_errors.RestErr) {

	pods, err := corev1.Pod.GetPods(ctx, client)
	if err != nil {
		return nil, rest_errors.NewBadRequestError(err.Error())
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
