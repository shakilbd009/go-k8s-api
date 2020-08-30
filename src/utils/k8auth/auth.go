package k8auth

import (
	"fmt"
	"log"
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	config = "kube_config"
)

var (
	kubeConfig = os.Getenv(config)
	Client     *kubernetes.Clientset
)

func init() {
	if kubeConfig == "" {
		log.Fatalln(fmt.Errorf("env variable %s is not set", config))
	}
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	if err != nil {
		log.Fatalln(err)
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln(err)
	}
	Client = client
}
