package services

import (
	"context"
	"strings"

	"github.com/shakilbd009/go-k8s-api/src/client/kubernetes/corev1"
	"github.com/shakilbd009/go-k8s-api/src/domain/k8s"
	"github.com/shakilbd009/go-k8s-api/src/utils/k8auth"
	"github.com/shakilbd009/go-k8s-api/src/utils/utility"
	"github.com/shakilbd009/go-utils-lib/rest_errors"
)

var (
	KDeploymentServices kDeploymentInterface = &kDeploymentService{}
	msg                                      = "no route to host"
)

type kDeploymentService struct{}

type kDeploymentInterface interface {
	CreateDeployment(context.Context, *k8s.K8sDeployment) (*k8s.K8sDeployment, rest_errors.RestErr)
	DeleteDeployment(context.Context, *k8s.K8sDeployment) (*k8s.K8sDeployment, rest_errors.RestErr)
}

func (*kDeploymentService) CreateDeployment(ctx context.Context, k *k8s.K8sDeployment) (*k8s.K8sDeployment, rest_errors.RestErr) {

	if err := k.ValidateCreateDeployment(); err != nil {
		return nil, err
	}
	_, err := AuthService.GetToken(k.Token)
	if err != nil {
		return nil, err
	}
	client := k8auth.Client
	if err := corev1.Namespace.Get(ctx, client, k.Namespace); err != nil {
		if strings.Contains(err.Error(), msg) {
			return nil, rest_errors.NewInternalServerError("Server unavailable", utility.ErrDatabase)
		}
		return nil, rest_errors.NewBadRequestError(err.Error())
	}

	result, err := k.Create(ctx, client)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (*kDeploymentService) DeleteDeployment(ctx context.Context, k *k8s.K8sDeployment) (*k8s.K8sDeployment, rest_errors.RestErr) {

	if err := k.ValidateDeleteDeployment(); err != nil {
		return nil, err
	}
	if _, err := AuthService.GetToken(k.Token); err != nil {
		return nil, err
	}
	client := k8auth.Client
	if err := corev1.Namespace.Get(ctx, client, k.Namespace); err != nil {
		return nil, rest_errors.NewBadRequestError(err.Error())
	}
	var resp k8s.K8sDeployment
	result, err := k.Delete(ctx, client)
	if err != nil {
		return nil, err
	}
	resp.DeploymentName = result.DeploymentName
	resp.Namespace = result.Namespace
	resp.Status = result.Status
	return &resp, nil
}
