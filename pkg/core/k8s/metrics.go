package k8s

import (
	"context"
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
//@function: PrometheusHealth
//@description: 检查Prometheus是否健康
//@param: clusterId uint
//@return:err error

func PrometheusHealth(clusterId uint) (prometheusUrl string, err error) {
	var clusterObj model.K8SCluster
	if err = options.DB.Where("id = ?", clusterId).First(&clusterObj).Error; err != nil {
		return "", err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()

	prourl := fmt.Sprintf("%s/-/ready", clusterObj.PrometheusUrl)
	readyReq, err := http.NewRequest("GET", prourl, nil)
	if err != nil {
		return "", err
	}

	readyResp, err := http.DefaultClient.Do(readyReq.WithContext(ctx))
	if err != nil {
		return "", err
	}

	// 如果还没有ready，则直接返回前端空数据
	if readyResp.StatusCode != http.StatusOK {
		return "", err
	}

	return clusterObj.PrometheusUrl, err
}

//@author: Eagle
//@function: GetMetricsData
//@description: 获取普罗米修斯数据
//@param: mt model.MetricsQuery
//@return: t map[string]*model.PrometheusQueryResp, err error

func GetMetricsData(mt model.MetricsQuery) (t map[string]*model.PrometheusQueryResp, err error) {
	step := 60
	end := time.Now().Unix()
	start := end - 3600

	// tracker
	tracker := model.NewPrometheusTracker()
	wg := sync.WaitGroup{}
	e := reflect.ValueOf(&mt).Elem()
	for i := 0; i < e.NumField(); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fName := e.Type().Field(i).Name
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
			prometheusUrl, err := PrometheusHealth(fValue.ClusterId)
			if fValue.Start != 0 || fValue.End != 0 {
				start = fValue.Start / 1000
				end = fValue.End / 1000
			}

			if err != nil {
				log.Printf("Prometheus health failed: " + err.Error())
				return
			}

			//http Get请求 Prometheus接口
			promql := url.QueryEscape(prometheusQueries.GetValueByField(fName))
			log.Printf("promql:" + prometheusQueries.GetValueByField(fName))
			fullpromql := fmt.Sprintf("%s/api/v1/query_range?query=%s&start=%d&end=%d&step=%d", prometheusUrl, promql, start, end, step)
			fmt.Println(fullpromql)
			resp, err := http.Get(fullpromql)
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

			var data model.PrometheusQueryResp
			if err := json.Unmarshal(body, &data); err != nil {
				log.Fatalf("unmarshal response body to models.PrometheusQueryResp failed" + err.Error())
				return
			}

			// 配置当前查询的数据结果
			tag := fTag.Get("json")
			tracker.Set(tag[:strings.Index(tag, ",omitempty")], &data)
		}(i)
	}

	// 等待所有查询完成
	wg.Wait()
	return tracker.Metrics, err
}
