package model

type K8SCluster struct {
	Id             int64  `gorm:"column:id;primary_key;AUTO_INCREMENT;not null" json:"id"`
	ApiAddress     string `gorm:"type:longText" json:"api_address" form:"api_address" gorm:"comment:api_address"`
	PrometheusUrl  string `gorm:"type:longText" json:"prometheus_url" form:"prometheus_url" gorm:"comment:prometheus_url"`
	PrometheusType int64  `gorm:"type:longText" json:"prometheus_type" form:"prometheus_type" gorm:"comment:prometheus_type"`
	PrometheusUser string `gorm:"type:longText" json:"prometheus_user" form:"prometheus_user" gorm:"comment:prometheus_user"`
	PrometheusPwd  string `gorm:"type:longText" json:"prometheus_pwd" form:"prometheus_pwd" gorm:"comment:prometheus_pwd"`
	ClusterName    string `gorm:"comment:集群名称;index:idx_cluster_name,unique" binding:"required" json:"clusterName"`
	KubeConfig     string `gorm:"comment:集群凭证;type:text" binding:"required" json:"kubeConfig"`
	NodeNumber     int    `gorm:"comment:节点数" json:"nodeNumber" `
}

func (k K8SCluster) TableName() string {
	return "k8s_cluster"
}
