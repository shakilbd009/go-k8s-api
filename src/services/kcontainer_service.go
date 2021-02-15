package services

import (
	"context"

	"github.com/shakilbd009/go-k8s-api/src/domain/k8s"
	"github.com/shakilbd009/go-k8s-api/src/utils/k8auth"
	"github.com/shakilbd009/go-utils-lib/rest_errors"
)

var KContainerServices kcontainerInterface = &kcontainerService{}

type kcontainerInterface interface {
	GetContainerLogs(ctx context.Context, namespace, podName, container string) ([]string, rest_errors.RestErr)
	GetContainers(ctx context.Context, namespace, podName string) ([]k8s.Kcontainer, rest_errors.RestErr)
}
type kcontainerService struct{}

func (*kcontainerService) GetContainers(ctx context.Context, namespace, podName string) ([]k8s.Kcontainer, rest_errors.RestErr) {
	var cont k8s.Kcontainer
	client := k8auth.Client
	containers, err := cont.GetContainers(ctx, client, namespace, podName)
	if err != nil {
		return nil, err
	}
	return containers, nil
}

func (*kcontainerService) GetContainerLogs(ctx context.Context, namespace, podName, container string) ([]string, rest_errors.RestErr) {
	client := k8auth.Client
	var cont k8s.Kcontainer
	logs, err := cont.GetContainerLogs(ctx, client, namespace, podName, container)
	if err != nil {
		return nil, err
	}
	return logs, nil
}
