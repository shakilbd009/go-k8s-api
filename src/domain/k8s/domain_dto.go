package k8s

import (
	"fmt"

	"github.com/shakilbd009/go-utils-lib/rest_errors"
)

var (
	msgTemplate     = "%v field is missing values"
	namespace       = "namespace"
	deployment_name = "deployment_name"
	container_name  = "container_name"
	image_version   = "image_version"
	replicas        = "replicas"
	required_field  = "required fields are missing"
)

type K8sDeployment struct {
	Namespace      string `json:"namespace,omitempty"`
	DeploymentName string `json:"deployment_name,omitempty"`
	ContainerName  string `json:"container_name,omitempty"`
	Image          string `json:"image_version,omitempty"`
	Replicas       int32  `json:"replicas,omitempty"`
	Status         string `json:"status,omitempty"`
	CreationTime   string `json:"creation_time,omitempty"`
}

func (k *K8sDeployment) ValidateCreateDeployment() rest_errors.RestErr {

	if k == nil {
		return rest_errors.NewBadRequestError(required_field)
	}
	if k.Namespace == "" {
		return rest_errors.NewBadRequestError(fmt.Sprintf(msgTemplate, namespace))
	}
	if k.DeploymentName == "" {
		return rest_errors.NewBadRequestError(fmt.Sprintf(msgTemplate, deployment_name))
	}
	if k.ContainerName == "" {
		return rest_errors.NewBadRequestError(fmt.Sprintf(msgTemplate, container_name))
	}
	if k.Image == "" {
		return rest_errors.NewBadRequestError(fmt.Sprintf(msgTemplate, image_version))
	}
	if k.Replicas == 0 {
		return rest_errors.NewBadRequestError(fmt.Sprintf(msgTemplate, replicas))
	}
	return nil
}

func (k *K8sDeployment) ValidateDeleteDeployment() rest_errors.RestErr {

	if k == nil {
		return rest_errors.NewBadRequestError(required_field)
	}
	if k.Namespace == "" {
		return rest_errors.NewBadRequestError(fmt.Sprintf(msgTemplate, namespace))
	}
	if k.DeploymentName == "" {
		return rest_errors.NewBadRequestError(fmt.Sprintf(msgTemplate, deployment_name))
	}
	return nil
}
