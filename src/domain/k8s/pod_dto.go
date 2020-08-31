package k8s

type Kpod struct {
	PodNamespace string `json:"pod_namespace"`
	PodName      string `json:"pod_name"`
	Status       string `json:"status"`
	Restarts     int32  `json:"restarts"`
	PodIP        string `json:"pod_ip"`
}
