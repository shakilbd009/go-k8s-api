package services

import (
	"context"

	"github.com/shakilbd009/go-k8s-api/src/client/kubernetes/corev1"
	"github.com/shakilbd009/go-k8s-api/src/domain/k8s"
	"github.com/shakilbd009/go-k8s-api/src/utils/k8auth"
	"github.com/shakilbd009/go-utils-lib/rest_errors"
)

var KPodServices kpodInterface = &kpodService{}

type kpodInterface interface {
	GetPods(ctx context.Context, namespace string, token string) ([]k8s.Kpod, rest_errors.RestErr)
}
type kpodService struct{}

func (*kpodService) GetPods(ctx context.Context, namespace, token string) ([]k8s.Kpod, rest_errors.RestErr) {

	if _, err := AuthService.GetToken(token); err != nil {
		return nil, err
	}
	client := k8auth.Client
	if err := corev1.Namespace.Get(ctx, client, namespace); err != nil {
		return nil, rest_errors.NewBadRequestError(err.Error())
	}
	var pod k8s.Kpod
	pods, err := pod.GetPods(ctx, client, namespace)
	if err != nil {
		return nil, err
	}
	return pods, nil
}
