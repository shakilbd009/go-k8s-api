package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shakilbd009/go-k8s-api/src/domain/k8s"
	"github.com/shakilbd009/go-k8s-api/src/services"
	"github.com/shakilbd009/go-k8s-api/src/utils/k8auth"
	"github.com/shakilbd009/go-utils-lib/rest_errors"
)

var (
	Kcontroller kcontrollerInterface = &kcontroller{}
	invalidJSON                      = "invalid json body"
)

type kcontroller struct{}

type kcontrollerInterface interface {
	CreateDeployment(*gin.Context)
	CreateMultiContainerDeployment(c *gin.Context)
	DeleteDeployment(*gin.Context)
	AddUser(*gin.Context)
	GetPods(*gin.Context)
	GetPodLogs(*gin.Context)
	GetContainers(c *gin.Context)
	GetContainerLogs(c *gin.Context)
}

func (k *kcontroller) CreateDeployment(c *gin.Context) {
	var deployment k8s.K8sDeployment
	if err := c.ShouldBindJSON(&deployment); err != nil {
		fmt.Println(err)
		restErr := rest_errors.NewBadRequestError(invalidJSON)
		c.JSON(restErr.Status(), restErr)
		return
	}
	result, err := services.KDeploymentServices.CreateDeployment(c.Request.Context(), &deployment)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, *result)
}

func (k *kcontroller) CreateMultiContainerDeployment(c *gin.Context) {
	var deployment k8s.K8sDeployments
	if err := c.ShouldBindJSON(&deployment); err != nil {
		fmt.Println(err)
		restErr := rest_errors.NewBadRequestError(invalidJSON)
		c.JSON(restErr.Status(), restErr)
		return
	}
	result, err := services.KDeploymentServices.CreateMultiContainerDeployment(c.Request.Context(), &deployment)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, *result)
}

func (k *kcontroller) DeleteDeployment(c *gin.Context) {

	var deployment k8s.K8sDeployment
	if err := c.ShouldBindJSON(&deployment); err != nil {
		restErr := rest_errors.NewBadRequestError(invalidJSON)
		c.JSON(restErr.Status(), restErr)
		return
	}

	result, err := services.KDeploymentServices.DeleteDeployment(c.Request.Context(), &deployment)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, *result)
}

func (k *kcontroller) GetPods(c *gin.Context) {

	token, ok := c.GetQuery("token")
	if !ok {
		rest_err := rest_errors.NewBadRequestError("token is missing as query param")
		c.JSON(rest_err.Status(), rest_err)
		return
	}
	namespace, ok := c.GetQuery("namespace")
	if !ok {
		rest_err := rest_errors.NewBadRequestError("namespace is missing as query param")
		c.JSON(rest_err.Status(), rest_err)
		return
	}
	result, err := services.KPodServices.GetPods(c.Request.Context(), namespace, token)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	if len(result) == 0 {
		restErr := rest_errors.NewNotFoundError(fmt.Sprintf("No resources found in %s namespace", namespace))
		c.JSON(restErr.Status(), restErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (k *kcontroller) GetPodLogs(c *gin.Context) {

	podName, ok := c.GetQuery("podName")
	if !ok {
		rest_err := rest_errors.NewBadRequestError("podName is missing as query param")
		c.JSON(rest_err.Status(), rest_err)
		return
	}
	namespace, ok := c.GetQuery("namespace")
	if !ok {
		rest_err := rest_errors.NewBadRequestError("namespace is missing as query param")
		c.JSON(rest_err.Status(), rest_err)
		return
	}
	logs, err := services.KPodServices.GetPodLogs(c.Request.Context(), namespace, podName)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, logs)
}

func (k *kcontroller) GetContainers(c *gin.Context) {

	podName, ok := c.GetQuery("podName")
	if !ok {
		rest_err := rest_errors.NewBadRequestError("podName is missing as query param")
		c.JSON(rest_err.Status(), rest_err)
		return
	}
	namespace, ok := c.GetQuery("namespace")
	if !ok {
		rest_err := rest_errors.NewBadRequestError("namespace is missing as query param")
		c.JSON(rest_err.Status(), rest_err)
		return
	}
	result, err := services.KContainerServices.GetContainers(c.Request.Context(), namespace, podName)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	if len(result) == 0 {
		restErr := rest_errors.NewNotFoundError(fmt.Sprintf("No resources found in %s namespace", namespace))
		c.JSON(restErr.Status(), restErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (k *kcontroller) GetContainerLogs(c *gin.Context) {

	podName, ok := c.GetQuery("podName")
	if !ok {
		rest_err := rest_errors.NewBadRequestError("podName is missing as query param")
		c.JSON(rest_err.Status(), rest_err)
		return
	}
	containerName, ok := c.GetQuery("containerName")
	if !ok {
		rest_err := rest_errors.NewBadRequestError("containerName is missing as query param")
		c.JSON(rest_err.Status(), rest_err)
		return
	}
	namespace, ok := c.GetQuery("namespace")
	if !ok {
		rest_err := rest_errors.NewBadRequestError("namespace is missing as query param")
		c.JSON(rest_err.Status(), rest_err)
		return
	}
	logs, err := services.KContainerServices.GetContainerLogs(c.Request.Context(), namespace, podName, containerName)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, logs)
}

func (k *kcontroller) AddUser(c *gin.Context) {

	var user k8s.K8sUser
	if err := c.ShouldBindJSON(&user); err != nil {
		rest_err := rest_errors.NewBadRequestError(invalidJSON)
		c.JSON(rest_err.Status(), rest_err)
		return
	}
	result, err := services.AuthService.AddUser(c.Request.Context(), k8auth.Client, &user)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, result)
}
