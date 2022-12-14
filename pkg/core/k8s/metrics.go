package k8s

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/dnsjia/luban/cmd/options"
	"github.com/dnsjia/luban/pkg/model"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"sync"
	"time"
)

// @Tags K8s_Metrics
// @Summary 获取Pod详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MetricsQuery true "获取Pod详情"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /kubernetes/metrics/getMetrics [post]
func GetMetrics(c *gin.Context) {
	var metricsQuery model.MetricsQuery
	_ = c.ShouldBindJSON(&metricsQuery)
	if queryResp, err := GetMetricsData(metricsQuery); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": "",
			"msg":  err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": queryResp,
			"msg":  "数据获取成功",
		})
		return
	}
}

//@author: Eagle
//@function: PrometheusAuth
//@description: Base Auth 认证信息
//@param: clusterObj model.K8SCluster
//@return:err error

func PrometheusAuth(clusterObj model.K8SCluster) (httpClient http.Client) {
	httpClient = http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
	if clusterObj.PrometheusType == 2 {
		httpClient = http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
				Proxy: func(req *http.Request) (*url.URL, error) {
					req.SetBasicAuth(clusterObj.PrometheusUser, clusterObj.PrometheusPwd)
					return nil, nil
				},
			}}
	}
	return httpClient
}

//@author: Eagle
//@function: PrometheusHealth
//@description: 检查Prometheus是否健康
//@param: clusterId uint
//@return:err error

func PrometheusHealth(clusterId uint) (prometheusUrl string, hclient http.Client, err error) {
	var clusterObj model.K8SCluster
	if err = options.DB.Where("id = ?", clusterId).First(&clusterObj).Error; err != nil {
		return "", hclient, err
	}

	// 超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()

	// 生成httpClient
	httpClient := PrometheusAuth(clusterObj)

	// 请求普罗米修斯ready接口
	prourl := fmt.Sprintf("%s/-/ready", clusterObj.PrometheusUrl)
	readyReq, err := http.NewRequest("GET", prourl, nil)
	if err != nil {
		return "", hclient, err
	}

	readyResp, err := httpClient.Do(readyReq.WithContext(ctx))
	if err != nil {
		return "", hclient, err
	}

	// 如果还没有ready，则直接返回前端空数据
	if readyResp.StatusCode != http.StatusOK {
		return "", hclient, err
	}

	return clusterObj.PrometheusUrl, httpClient, err
}

//@author: Eagle
//@function: GetMetricsData
//@description: 获取普罗米修斯数据
//@param: mt model.MetricsQuery
//@return: t map[string]*model.PrometheusQueryResp, err error

func GetMetricsData(mt model.MetricsQuery) (t map[string]*model.PrometheusQueryResp, err error) {
	// 查询时间范围
	step := 60
	end := time.Now().Unix()
	start := end - 3600

	// 初始化tracker
	tracker := model.NewPrometheusTracker()

	// 声明一个等待组
	wg := sync.WaitGroup{}

	// 通过反射获信息
	e := reflect.ValueOf(&mt).Elem()

	for i := 0; i < e.NumField(); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			// 字段名称，例如：  MemoryUsage, PodUsage
			fName := e.Type().Field(i).Name

			// 断言结构体值填充到结构体赋值
			fValue := e.Field(i).Interface().(*model.MetricsCategory)
			fTag := e.Type().Field(i).Tag
			if fValue == nil {
				return
			}

			// 请求 Prometheus 查询
			prometheusQueries := fValue.GenerateQuery()
			if prometheusQueries == nil {
				log.Printf("prometheusQueries nill failed" + err.Error())
				return
			}

			//Prometheus健康检查
			prometheusUrl, httpClient, err := PrometheusHealth(fValue.ClusterId)
			if fValue.Start != 0 || fValue.End != 0 {
				start = fValue.Start / 1000
				end = fValue.End / 1000
			}

			if err != nil {
				log.Printf("Prometheus health failed: " + err.Error())
				return
			}

			// 通过反射返回promsql
			promql := url.QueryEscape(prometheusQueries.GetValueByField(fName))
			log.Printf("promsql 打印:" + prometheusQueries.GetValueByField(fName))

			fullpromql := fmt.Sprintf("%s/api/v1/query_range?query=%s&start=%d&end=%d&step=%d", prometheusUrl, promql, start, end, step)
			//http Get请求 Prometheus接口
			resp, err := httpClient.Get(fullpromql)
			if err != nil {
				log.Fatalf("request metrics data failed" + err.Error())
				return
			}

			//Prometheus 接口返回数据处理
			defer func(Body io.ReadCloser) {
				_ = Body.Close()
			}(resp.Body)

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Printf("read response body failed" + err.Error())
				return
			}

			// 解析数据到结构体
			var data model.PrometheusQueryResp
			if err := json.Unmarshal(body, &data); err != nil {
				log.Fatalf("unmarshal response body to models.PrometheusQueryResp failed" + err.Error())
				return
			}

			// 从model tag里面获取里面数据
			tag := fTag.Get("json")

			// json:"memoryUsage,omitempty" 获取里面memoryUsage作为key
			tracker.Set(tag[:strings.Index(tag, ",omitempty")], &data)
		}(i)
	}

	// 等待所有查询完成
	wg.Wait()
	return tracker.Metrics, err
}
