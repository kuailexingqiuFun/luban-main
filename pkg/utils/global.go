package utils

import (
	"github.com/dnsjia/luban/cmd/options"
	"github.com/dnsjia/luban/pkg/model"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type kubeclient struct {
	Pod *Pod
}

func NewKubeClient(id int) (kube *kubeclient, err error) {
	var cluster model.K8SCluster
	if err := options.DB.Where("id = ?", id).First(&cluster).Error; err != nil {
		return nil, err
	}

	var config *rest.Config
	config, err = clientcmd.RESTConfigFromKubeConfig([]byte(cluster.KubeConfig))
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	client := &kubeclient{
		Pod: NewPod(clientset, config),
	}

	return client, nil
}
