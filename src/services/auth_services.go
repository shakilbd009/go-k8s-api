package services

import (
	"context"

	"github.com/shakilbd009/go-k8s-api/src/auth/data"
	"github.com/shakilbd009/go-k8s-api/src/client/kubernetes/corev1"
	"github.com/shakilbd009/go-k8s-api/src/domain/k8s"
	"github.com/shakilbd009/go-utils-lib/rest_errors"
	"k8s.io/client-go/kubernetes"
)

var (
	AuthService  authInterface = &authService{}
	doesNotExist               = "namespace does not exist"
)

type authService struct{}
type authInterface interface {
	AddUser(context.Context, *kubernetes.Clientset, *k8s.K8sUser) (*k8s.K8sUser, rest_errors.RestErr)
	GetToken(string) (*k8s.K8sUser, rest_errors.RestErr)
}

func (*authService) AddUser(ctx context.Context, client *kubernetes.Clientset, u *k8s.K8sUser) (*k8s.K8sUser, rest_errors.RestErr) {

	if err := u.Validate(); err != nil {
		return nil, err
	}
	if err := corev1.Namespace.Get(ctx, client, u.Namespace); err != nil {
		return nil, rest_errors.NewNotFoundError(err.Error())
	}
	user, err := data.Database.CreateUser(u)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (*authService) GetToken(token string) (*k8s.K8sUser, rest_errors.RestErr) {
	return data.Database.GetAccessToken(token)
}
