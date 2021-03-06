package corev1

import (
	"bytes"
	"context"
	"io"
	"strings"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var (
	Pod podIinterface = &pod{}
)

type podIinterface interface {
	GetPods(context.Context, *kubernetes.Clientset, string) (*corev1.PodList, error)
	GetPodLogs(ctx context.Context, client *kubernetes.Clientset, namespace, podName string) ([]string, error)
	GetContainerLogs(ctx context.Context, client *kubernetes.Clientset, namespace, podName, container string) ([]string, error)
	GetContainers(ctx context.Context, client *kubernetes.Clientset, namespace, podName string) ([]corev1.Container, error)
}
type pod struct{}

func (*pod) GetPods(ctx context.Context, client *kubernetes.Clientset, namespace string) (*corev1.PodList, error) {

	pods, err := client.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return pods, nil
}

func (*pod) GetContainers(ctx context.Context, client *kubernetes.Clientset, namespace, podName string) ([]corev1.Container, error) {

	pod, err := client.CoreV1().Pods(namespace).Get(ctx, podName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return pod.Spec.Containers, nil
}

func (*pod) GetPodLogs(ctx context.Context, client *kubernetes.Clientset, namespace, podName string) ([]string, error) {

	lines := int64(5)
	request := client.CoreV1().Pods(namespace).GetLogs(podName, &corev1.PodLogOptions{
		TailLines: &lines,
	})
	stream, err := request.Stream(ctx)
	if err != nil {
		return nil, err
	}
	defer stream.Close()
	buf := &bytes.Buffer{}
	_, err = io.Copy(buf, stream)
	if err != nil {
		return nil, err
	}
	logs := strings.Split(buf.String(), "\n")
	return logs, nil
}

func (*pod) GetContainerLogs(ctx context.Context, client *kubernetes.Clientset, namespace, podName, container string) ([]string, error) {

	lines := int64(5)
	request := client.CoreV1().Pods(namespace).GetLogs(podName, &corev1.PodLogOptions{
		TailLines: &lines,
		Container: container,
	})
	stream, err := request.Stream(ctx)
	if err != nil {
		return nil, err
	}
	defer stream.Close()
	buf := &bytes.Buffer{}
	_, err = io.Copy(buf, stream)
	if err != nil {
		return nil, err
	}
	logs := strings.Split(buf.String(), "\n")
	return logs, nil
}
