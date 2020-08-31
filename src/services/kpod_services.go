package services

import (
	"context"

	"github.com/shakilbd009/go-k8s-api/src/domain/k8s"
	"github.com/shakilbd009/go-k8s-api/src/utils/k8auth"
	"github.com/shakilbd009/go-utils-lib/rest_errors"
)

var KPodServices kpodInterface = &kpodService{}

type kpodInterface interface {
	GetPods(context.Context) ([]k8s.Kpod, rest_errors.RestErr)
}
type kpodService struct{}

func (*kpodService) GetPods(ctx context.Context) ([]k8s.Kpod, rest_errors.RestErr) {

	var pod k8s.Kpod
	pods, err := pod.GetPods(ctx, k8auth.Client)
	if err != nil {
		return nil, err
	}
	return pods, nil
}
