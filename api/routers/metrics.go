package routers

import (
	"github.com/dnsjia/luban/pkg/core/k8s"
	"github.com/gin-gonic/gin"
)

func MetricsApiRouter(r *gin.RouterGroup) {
	MetricsRouter := r.Group("kubernetes/metrics")
	MetricsRouter.POST("/getMetrics", k8s.GetMetrics) // 监控数据
}
