package temp

import (
	"context"
	"fmt"
	"log"

	"github.com/shakilbd009/go-k8s-api/src/client/kubernetes/appsv1"
	"github.com/shakilbd009/go-k8s-api/src/utils/utility"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func delete() {

	config, err := clientcmd.BuildConfigFromFlags("", "/home/shakil/.kube/config")
	if err != nil {
		log.Fatalln(err)
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln(err)
	}
	if err := appsv1.Deployment.Delete(context.Background(), client, "test-nginx", "test-nginx-demo"); err != nil {
		log.Fatalln(err)
	}
}

func createDeployment(client *kubernetes.Clientset) {
	resp, err := appsv1.Deployment.Create(
		context.Background(),
		client,
		"test-nginx",
		"test-nginx-demo",
		"nginx",
		"nginx:1.14.2",
		utility.GetInt32Pointer(3),
	)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp)
}

func test() {
	// nodes, err := client.CoreV1().Events("default").List(context.Background(), v1.ListOptions{})
	// pods, err := client.CoreV1().Pods(v1.NamespaceAll).List(context.Background(), v1.ListOptions{})
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// //btyes := make([]byte, 5201)
	// fmt.Printf("cluster-name: %-6s name: %-36s status: %10s restarts: %s", "", "", "", "")
	// fmt.Println()
	// for _, pod := range pods.Items {
	// 	fmt.Println(fmt.Sprintf("%-20v %-42v %10s %v", pod.Namespace, pod.Name, pod.Status.Phase, pod.Status.ContainerStatuses[0].RestartCount))
	// }

	//nodes, err := client.CoreV1().Events("default").List(context.Background(), v1.ListOptions{})
	// pods, err := client.CoreV1().Pods("argocd").List(v2.ListOptions{})
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Printf("Namespace: %-9s pod-name: %-37s pod-status: %-2s restarts: %s podIP: %-5s", "", "", "", "", "")
	// fmt.Println()
	// for _, pod := range pods.Items {
	// 	fmt.Println(fmt.Sprintf("%-20v %-47v %-14s %-10v %-5s", pod.Namespace, pod.Name, pod.Status.Phase, pod.Status.ContainerStatuses[0].RestartCount, pod.Status.PodIP))
	// }
}
