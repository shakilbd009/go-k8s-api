package services

import (
	"context"

	"github.com/shakilbd009/go-k8s-api/src/domain/k8s"
	"github.com/shakilbd009/go-k8s-api/src/utils/k8auth"
	"github.com/shakilbd009/go-utils-lib/rest_errors"
)

var Kservice kserviceInterface = &kservice{}

type kservice struct{}

type kserviceInterface interface {
	CreateDeployment(context.Context, *k8s.K8sDeployment) (*k8s.K8sDeployment, rest_errors.RestErr)
	DeleteDeployment(context.Context, *k8s.K8sDeployment) (*k8s.K8sDeployment, rest_errors.RestErr)
}

func (*kservice) CreateDeployment(ctx context.Context, k *k8s.K8sDeployment) (*k8s.K8sDeployment, rest_errors.RestErr) {

	if err := k.ValidateCreateDeployment(); err != nil {
		return nil, err
	}
	result, err := k.Create(ctx, k8auth.Client)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (*kservice) DeleteDeployment(ctx context.Context, k *k8s.K8sDeployment) (*k8s.K8sDeployment, rest_errors.RestErr) {

	if err := k.ValidateDeleteDeployment(); err != nil {
		return nil, err
	}
	result, err := k.Delete(ctx, k8auth.Client)
	if err != nil {
		return nil, err
	}
	return result, nil
}
