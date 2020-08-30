package appsv1

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const (
	appName = "app"
	http    = "http"
	port80  = 80
)

var (
	Deployment deploymentInterface = &deployment{}
)

type deploymentInterface interface {
	Create(ctx context.Context, client *kubernetes.Clientset, namespace, deploymentName, containerName, image string, replicas *int32) (string, error)
	Delete(ctx context.Context, client *kubernetes.Clientset, namespace, deploymentName string) error
}

type deployment struct{}

func (*deployment) Create(ctx context.Context, client *kubernetes.Clientset, namespace, deploymentName, containerName, image string, replicas *int32) (string, error) {

	param := appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: deploymentName,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					appName: containerName,
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						appName: containerName,
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  containerName,
							Image: image,
							Ports: []apiv1.ContainerPort{
								{
									Name:          http,
									ContainerPort: port80,
								},
							},
						},
					},
				},
			},
		},
	}
	result, err := client.AppsV1().Deployments(namespace).Create(ctx, &param, metav1.CreateOptions{})
	if err != nil {
		return "", err
	}
	return result.CreationTimestamp.String(), nil
	//return result.Status.String(), nil
}

func (*deployment) Delete(ctx context.Context, client *kubernetes.Clientset, namespace, deploymentName string) error {
	return client.AppsV1().Deployments(namespace).Delete(ctx, deploymentName, metav1.DeleteOptions{})
}
