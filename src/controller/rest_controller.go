package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shakilbd009/go-k8s-api/src/domain/k8s"
	"github.com/shakilbd009/go-k8s-api/src/services"
	"github.com/shakilbd009/go-utils-lib/rest_errors"
)

var Kcontroller kcontrollerInterface = &kcontroller{}

type kcontroller struct{}

type kcontrollerInterface interface {
	CreateDeployment(*gin.Context)
	DeleteDeployment(*gin.Context)
}

func (k *kcontroller) CreateDeployment(c *gin.Context) {
	var deployment k8s.K8sDeployment
	if err := c.ShouldBindJSON(&deployment); err != nil {
		fmt.Println(err)
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}
	result, err := services.Kservice.CreateDeployment(c.Request.Context(), &deployment)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, *result)
}

func (k *kcontroller) DeleteDeployment(c *gin.Context) {

	var deployment k8s.K8sDeployment
	if err := c.ShouldBindJSON(&deployment); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}
	result, err := services.Kservice.DeleteDeployment(c.Request.Context(), &deployment)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, *result)
}
