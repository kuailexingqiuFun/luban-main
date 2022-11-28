package k8s

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/dnsjia/luban/api/types"
	"github.com/dnsjia/luban/cmd/options"
	"github.com/dnsjia/luban/pkg/model"
)

func CreateCluster(c *gin.Context) {
	var (
		clusterReq types.ClusterRequest
		cluster    model.K8SCluster
	)
	if err := c.ShouldBindJSON(&clusterReq); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": "",
			"msg":  err.Error(),
		})
		return
	}

	// 创建clientSet
	clientSet, err := GetClientSet(clusterReq.KubeConfig)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": "",
			"msg":  err.Error(),
		})
		return
	}

	// 获取node数量
	nodeList, err := clientSet.CoreV1().Nodes().List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": "",
			"msg":  err.Error(),
		})
		return
	}

	// 插入数据
	if err := options.DB.Model(&model.K8SCluster{}).Where("cluster_name = ?", clusterReq.ClusterName).First(&cluster).Error; err == gorm.ErrRecordNotFound {
		if err := options.DB.Model(&model.K8SCluster{}).Create(&model.K8SCluster{
			ClusterName:    clusterReq.ClusterName,
			KubeConfig:     clusterReq.KubeConfig,
			ApiAddress:     clusterReq.ApiAddress,
			PrometheusType: clusterReq.PrometheusType,
			PrometheusUrl:  clusterReq.PrometheusUrl,
			PrometheusUser: clusterReq.PrometheusUser,
			PrometheusPwd:  clusterReq.PrometheusPwd,
			NodeNumber:     len(nodeList.Items),
		}).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"data": "",
				"msg":  err.Error(),
			})
			return
		}
	} else {
		if err := options.DB.Model(&model.K8SCluster{}).Where("id = ?", cluster.Id).Updates(
			map[string]interface{}{
				"kube_config":     clusterReq.KubeConfig,
				"api_address":     clusterReq.ApiAddress,
				"prometheus_url":  clusterReq.PrometheusUrl,
				"prometheus_type": clusterReq.PrometheusType,
				"prometheus_user": clusterReq.PrometheusUser,
				"prometheus_pwd":  clusterReq.PrometheusPwd,
			}).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"data": "",
				"msg":  err,
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": "",
		"msg":  "创建集群成功",
	})
}

func ListCluster(c *gin.Context) {
	var cluster []model.K8SCluster
	if err := options.DB.Model(&model.K8SCluster{}).Find(&cluster).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": "",
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": cluster,
		"msg":  "创建集群成功",
	})
}

func DeleteCluster(c *gin.Context) {
	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": "",
			"msg":  err.Error(),
		})
		return
	}

	if err := options.DB.Model(&model.K8SCluster{}).Where("id = ?", id).Delete(&model.K8SCluster{Id: id}).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": "",
			"msg":  err.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": "",
		"msg":  "操作成功",
	})

}

func UpdateCluster(c *gin.Context) {
	var cluster model.K8SCluster
	if err := c.ShouldBindJSON(&cluster); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": "",
			"msg":  err.Error(),
		})
		return
	}

	// 只更新集群kubeconfig字段
	if err := options.DB.Model(&model.K8SCluster{}).Where("id = ?", cluster.Id).Updates(
		map[string]interface{}{
			"kube_config": cluster.KubeConfig,
		}).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": "",
			"msg":  err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": "",
		"msg":  "操作成功",
	})
}

type ClusterOptions struct {
	Id string `uri:"id" binding:"required"`
}

func GetClusterDetail(c *gin.Context) {
	var cluster ClusterOptions
	if err := c.ShouldBindUri(&cluster); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": "",
			"msg":  err,
		})
		return
	}

	var k8s model.K8SCluster
	if err := options.DB.Model(&model.K8SCluster{}).Where("id = ?", cluster.Id).Find(&k8s).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": "",
			"msg":  err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": k8s,
		"msg":  "ok",
	})
}
