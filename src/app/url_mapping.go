package app

import "github.com/shakilbd009/go-k8s-api/src/controller"

func urlMapping() {
	router.POST("/deployment", controller.Kcontroller.CreateDeployment)
	router.POST("/mcdeployment", controller.Kcontroller.CreateMultiContainerDeployment)
	router.DELETE("/deployment", controller.Kcontroller.DeleteDeployment)
	router.GET("/pods", controller.Kcontroller.GetPods)
	router.GET("/podLogs", controller.Kcontroller.GetPodLogs)
	router.GET("/containers", controller.Kcontroller.GetContainers)
	router.GET("/containerLogs", controller.Kcontroller.GetContainerLogs)
	router.POST("/register", controller.Kcontroller.AddUser)
}
