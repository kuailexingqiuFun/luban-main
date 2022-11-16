package routers

import (
	"github.com/dnsjia/luban/pkg/core/k8s"
	"github.com/gin-gonic/gin"
)

func ProxyApiRouter(r *gin.RouterGroup) {
	ProxyRouter := r.Group("kubernetes/proxy")
	ProxyRouter.Any("clusters", k8s.ListCluster)
	ProxyRouter.Any("/:cluster_id/*path", k8s.K8sAPIProxy)
}
