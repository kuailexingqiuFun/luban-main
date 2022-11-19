package main

import (
	"github.com/gin-gonic/gin"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	Host := "https://127.0.0.1:50101"
	BearerToken := "eyJhbGciOiJSUzI1NiIsImtpZCI6IjAxNC04cUp5bElBRGVpS1JVY1ZsVVBQY3dRcHJfRnYwaUpqNVlRdTBKYlUifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJkZWZhdWx0Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZWNyZXQubmFtZSI6ImFkbWluLXVzZXItdG9rZW4iLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC5uYW1lIjoiYWRtaW4tdXNlciIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50LnVpZCI6IjRlYWRiMGI4LWM3NjctNDYzNy1hMjJkLTI1ZTY2MDc4NzJlNyIsInN1YiI6InN5c3RlbTpzZXJ2aWNlYWNjb3VudDpkZWZhdWx0OmFkbWluLXVzZXIifQ.OxIPZWmgaBee8l8fMsiS1vDdsAYKkzigkk5-8XVuuEm_PLWZ7nGCMT9DkzN2VjdeSXHyJkU6AinxoZsb1le8he0tq5ANQwQtCr135ZhvfB0EsqZpJDwu21kZ-IIXX5BTJKhL0y_ndJEi6SRmVxHRoWvdUGEFgEVUxxLj5gJMf-MGUjuxxSMIIiu8PlPbpn2yFKVYAnkLghN9YU-XhWfUgxcPQ2FfY8Q4pV9jaJ7dze4dNLPTZbvTK1MbnzDPmCo5d_oLvvGFwGriuw3FwY-pGn5MBYS1XNd4hZYP1s39x7U0rC_Yvaqd0TLYMdgeYxr-7ZTm3MZedXIh0DYX3oEBjw"
	config := &rest.Config{
		Host:            Host,
		BearerToken:     BearerToken,
		TLSClientConfig: rest.TLSClientConfig{Insecure: true},
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	sharedInformerFactory := informers.NewSharedInformerFactory(clientset, 0)
	stopCh := make(chan struct{})
	defer close(stopCh)

	engin := gin.Default()
	// 构造一个通用的根据不同group/version/kind 获取对象类型
	//http://localhost:8888/core/v1/pods
	engin.GET("/:group/:version/:resource", func(c *gin.Context) {
		g, v, r := c.Param("group"), c.Param("version"), c.Param("resource")
		// 如果是core组则为空字符串
		if g == "core" {
			g = ""
		}
		gvr := schema.GroupVersionResource{g, v, r}
		genericInformer, _ := sharedInformerFactory.ForResource(gvr)
		_ = genericInformer.Informer()
		sharedInformerFactory.Start(stopCh)
		sharedInformerFactory.WaitForCacheSync(stopCh)
		items, err := genericInformer.Lister().List(labels.Everything())
		if err != nil {
			c.JSON(500, gin.H{
				"msg": err,
			})
		}
		c.JSON(200, gin.H{
			"msg": items,
		})

	})
	engin.Run(":8080")
}
