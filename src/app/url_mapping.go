package app

import "github.com/shakilbd009/go-k8s-api/src/controller"

func urlMapping() {
	router.POST("/deployment", controller.Kcontroller.CreateDeployment)
	router.DELETE("/deployment", controller.Kcontroller.DeleteDeployment)
}
