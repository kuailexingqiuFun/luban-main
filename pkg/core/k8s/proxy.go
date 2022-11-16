package k8s

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dnsjia/luban/cmd/options"
	"github.com/dnsjia/luban/pkg/model"
	"github.com/dnsjia/luban/pkg/utils"
	"github.com/dnsjia/luban/pkg/utils/kubernetes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
	"net/url"
	"strings"
)

type ItemList []interface{}

type Item struct {
	Metadata metav1.ObjectMeta `json:"metadata"`
}

type K8sObj interface{}

type ProxyParamRequest struct {
	ClusterId int    `uri:"cluster_id" binding:"required"`
	Path      string `uri:"path" binding:"required"`
}

type ProxyRequest struct {
	Search        string `form:"search" binding:"required"`
	Namespace     string `form:"namespace" binding:"required"`
	Keywords      string `form:"keywords" binding:"required"`
	FieldSelector string `form:"fieldSelector" binding:"required"`
	LabelSelector string `form:"labelSelector" binding:"required"`
	Page          int    `form:"page" binding:"required"`
	PageSize      int    `form:"pageSize" binding:"required"`
}

type TerminalRequest struct {
	Name      string `json:"name"  form:"name"`
	PodName   string `json:"pod_name" form:"pod_name"`
	Namespace string `json:"namespace" form:"namespace"`
	ClusterId int    `json:"cluster_id" form:"cluster_id"`
	XToken    string `json:"x-token" form:"x-token"`
	Cols      int    `json:"cols" form:"cols"`
	Rows      int    `json:"rows" form:"rows"`
}

// @Tags ProxyApi
// @Summary 分页获取ProxyApi列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Router  /kubernetes/proxy/:cluster_id/*path [Any]

func K8sAPIProxy(c *gin.Context) {
	// 解析参数
	var proxy ProxyRequest
	_ = c.ShouldBindQuery(&proxy)
	var urlParam ProxyParamRequest
	_ = c.BindUri(&urlParam)

	ret, err := ProxyOption(c, proxy, urlParam)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": "",
			"msg":  err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": ret,
			"msg":  "数据获取成功",
		})
		return
	}

}

// 集群版本判断
func compatibleClusterVersion(minor int, path *string) {
	p := *path
	if minor <= 18 {
		if strings.Contains(p, "networking.k8s.io/v1") && strings.Contains(p, "ingresses") {
			p = strings.Replace(p, "networking.k8s.io/v1", "networking.k8s.io/v1beta1", -1)
		}
	}
	*path = p
}

// 参数拆分
func parseResourceName(path string) (string, error) {
	ss := strings.Split(path, "/")
	if len(ss) > 0 {
		return ss[len(ss)-1], nil
	}
	return "", fmt.Errorf("cant not get resource name from url %s", path)
}

// 选择器拼接
func Selector(proxy ProxyRequest, urlParam ProxyParamRequest) (ProxyRequest, ProxyParamRequest) {

	if proxy.FieldSelector != "" && proxy.LabelSelector != "" {
		urlParam.Path += "?fieldSelector=" + proxy.FieldSelector + "&labelSelector=" + proxy.LabelSelector
	} else if proxy.FieldSelector != "" && proxy.LabelSelector == "" {
		urlParam.Path += "?fieldSelector=" + proxy.FieldSelector
	} else if proxy.LabelSelector != "" {
		urlParam.Path += "?labelSelector=" + proxy.LabelSelector
	}

	return proxy, urlParam
}

// 多命名空间资源获取
func fetchMultiNamespace(c *gin.Context, proxy ProxyRequest, httpClient http.Client, apiUrl *url.URL, allowedNamespaces []string) (ret Result, err error) {
	//获取已有权限的命名空间资源
	resp, err := kubernetes.FetchMultiNamespaceResource(&httpClient, allowedNamespaces, *apiUrl)
	if err != nil {
		return ret, errors.New("fetchMultiNamespaceResource request failed: " + err.Error())
	}

	//分页
	PageResult, err := pagerAndSearch(proxy.Page, proxy.PageSize, resp.Items, proxy.Keywords)
	if err != nil {
		return ret, err
	}
	return *PageResult, err
}

// @Tags ProxyApi
// @Summary ProxyOption操作
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json

func ProxyOption(c *gin.Context, proxy ProxyRequest, urlParam ProxyParamRequest) (ret Result, err error) {
	var cluster model.K8SCluster
	if err = options.DB.Model(&model.K8SCluster{}).Where("id = ?", urlParam.ClusterId).First(&cluster).Error; err != nil {
		return ret, errors.New("select Cluster failed: " + err.Error())
	}

	search := false
	if proxy.Search != "" {
		search = true
	}

	// 生成transport
	ts, err := kubernetes.GenerateTLSTransport(&cluster)
	if err != nil {
		return ret, errors.New("Transport fail" + err.Error())
	}

	// 生成httpClient
	httpClient := http.Client{Transport: ts}
	k := kubernetes.NewKubernetes(&cluster)
	clusterVersionMinor, err := k.VersionMinor()
	if err != nil {
		return ret, errors.New("Minor Version cluster failed: " + err.Error())
	}

	compatibleClusterVersion(clusterVersionMinor, &urlParam.Path)

	//判断是否已经包含了namespace的查询
	resourceName, err := parseResourceName(strings.Split(urlParam.Path, "?")[0])
	if err != nil {
		return ret, errors.New("Get  resourceName failed: " + err.Error())
	}

	//判断资源类型是否是namespace级别的
	namespaced, err := k.IsNamespacedResource(resourceName)
	if err != nil {
		return ret, errors.New("IsNamespacedResource failed: " + err.Error())
	}

	// 判断是否包含监控接口，或者包含命名空间
	if strings.Contains(urlParam.Path, "namespaces") || strings.Contains(urlParam.Path, "metrics.k8s.io") {
		namespaced = false
	}

	//选择器拼接
	proxy, urlParam = Selector(proxy, urlParam)

	//url格式化
	apiUrl, err := url.Parse(fmt.Sprintf("%s%s", cluster.ApiAddress, urlParam.Path))
	if err != nil {
		return ret, errors.New("url Parse failed: " + err.Error())
	}

	// 调用多namespace逻辑
	if http.MethodGet == c.Request.Method && proxy.Namespace == "" && namespaced {
		//已获权的命名空间
		allowedNamespaces, err := k.GetUserNamespaceNames(c)
		if err != nil {
			return ret, err
		}

		return fetchMultiNamespace(c, proxy, httpClient, apiUrl, allowedNamespaces)
	}

	//创建http请求 调用单namespace逻辑
	req, err := http.NewRequest(c.Request.Method, apiUrl.String(), c.Request.Body)
	if err != nil {
		return ret, errors.New("http request failed: " + err.Error())
	}

	//修改头信息
	if c.Request.Method == "PATCH" {
		req.Header.Set("Content-Type", "application/merge-patch+json")
	}

	//发起请求
	resp, err := httpClient.Do(req)
	if err != nil {
		return ret, errors.New("http request do failed: " + err.Error())
	}

	//取出数据
	rawResp, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode == http.StatusForbidden {
		resp.StatusCode = http.StatusInternalServerError
	}

	//解析出数据
	if req.Method == http.MethodGet && search {
		var listObj K8sListObj
		if err = json.Unmarshal(rawResp, &listObj); err != nil {
			return ret, errors.New("json Unmarshal failed: " + string(rawResp))
		}

		//分页
		PageResult, err := pagerAndSearch(proxy.Page, proxy.PageSize, listObj.Items, proxy.Keywords)
		if err != nil {
			return ret, errors.New("pagerAndSearch  failed: " + err.Error())
		}

		return *PageResult, err
	}

	//解析数据
	if req.Method == http.MethodGet || req.Method == http.MethodPut || req.Method == http.MethodPost || req.Method == http.MethodPatch || req.Method == http.MethodDelete {
		var k8sObj K8sObj
		if err = json.Unmarshal(rawResp, &k8sObj); err != nil {
			return ret, errors.New("json Unmarshal failed: " + string(rawResp))
		}

		// 单个数据返回
		var PageResult Result
		PageResult.Items = k8sObj
		return PageResult, err
	}

	return ret, err
}

// @Tags WsApi
// @Summary  终端
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body kubernetesReq.TerminalRequest true "根据id获取api"
// @Success 200 {object} response.Response{data=systemRes.SysAPIResponse} "根据id获取api,返回包括api详情"
// @Router  /kubernetes/pods/terminal [get]
func Terminal(c *gin.Context) {
	var terminal TerminalRequest
	_ = c.ShouldBindQuery(&terminal)

	client, err := utils.NewKubeClient(terminal.ClusterId)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	var cluster model.K8SCluster
	if err = options.DB.Model(&model.K8SCluster{}).Where("id = ?", terminal.ClusterId).First(&cluster).Error; err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	//校验pod
	kubeshell, err := utils.NewKubeShell(c.Writer, c.Request, nil)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	cmd := []string{
		"/bin/sh", "-c", fmt.Sprintf("clear;(bash || sh); export LINES=%d ; export COLUMNS=%d;", terminal.Rows, terminal.Cols),
	}
	if err := client.Pod.Exec(cmd, kubeshell, terminal.Namespace, terminal.PodName, terminal.Name); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
}

// @Tags WsApi
// @Summary  终端日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body kubernetesReq.TerminalRequest true "根据id获取api"
// @Success 200 {object} response.Response{data=systemRes.SysAPIResponse} "根据id获取api,返回包括api详情"
// @Router  /kubernetes/pods/logs [get]
func ContainerLog(c *gin.Context) {
	var terminal TerminalRequest
	_ = c.ShouldBindQuery(&terminal)

	kubeLogger, err := utils.NewKubeLogger(c.Writer, c.Request, nil)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	client, err := utils.NewKubeClient(terminal.ClusterId)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	client.Pod.ContainerLog(kubeLogger, terminal.Name, terminal.PodName, terminal.Namespace)
}
