package kubernetes

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dnsjia/luban/pkg/model"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/version"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type Interface interface {
	Config() (*rest.Config, error)
	Client() (*kubernetes.Clientset, error)
	VersionMinor() (int, error)
	IsNamespacedResource(resourceName string) (bool, error)
	GetUserNamespaceNames(c *gin.Context) ([]string, error)
}

type Kubernetes struct {
	*model.K8SCluster
}

func NewKubernetes(cluster *model.K8SCluster) Interface {
	return &Kubernetes{cluster}
}

func (k *Kubernetes) Config() (config *rest.Config, err error) {
	return clientcmd.RESTConfigFromKubeConfig([]byte(k.KubeConfig))
}

func (k *Kubernetes) VersionMinor() (int, error) {
	v, err := k.Version()
	if err != nil {
		return 0, err
	}
	reg := regexp.MustCompile("[^0-9]")
	minor, err := strconv.Atoi(reg.ReplaceAllString(v.Minor, ""))
	return minor, err
}

func (k *Kubernetes) Version() (*version.Info, error) {
	client, err := k.Client()
	if err != nil {
		return nil, err
	}
	return client.ServerVersion()
}

func (k *Kubernetes) Client() (*kubernetes.Clientset, error) {
	cfg, err := k.Config()
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(cfg)
}

func (k *Kubernetes) GetUserNamespaceNames(c *gin.Context) ([]string, error) {
	client, err := k.Client()
	if err != nil {
		return nil, err
	}

	ns, err := client.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var nsList []string
	for i := range ns.Items {
		if ns.Items[i].Status.Phase == "Active" {
			nsList = append(nsList, ns.Items[i].Name)
		}
	}

	sort.Strings(nsList)
	return nsList, nil
}

func (k *Kubernetes) IsNamespacedResource(resourceName string) (bool, error) {
	if resourceName == "events" || resourceName == "nodes" {
		return false, nil
	}
	client, err := k.Client()
	if err != nil {
		return false, err
	}
	apiList, err := client.ServerPreferredNamespacedResources()
	if err != nil && len(apiList) == 0 {
		return false, err
	}
	for i := range apiList {
		fmt.Println("命名空间级别资源", apiList[i].APIResources)
		for j := range apiList[i].APIResources {
			if apiList[i].APIResources[j].Name == resourceName {
				return true, nil
			}
		}
	}
	return false, nil
}

func GenerateTLSTransport(c *model.K8SCluster) (http.RoundTripper, error) {
	kube := NewKubernetes(c)
	kubeconfig, err := kube.Config()
	if err != nil {
		return nil, err
	}

	return rest.TransportFor(kubeconfig)
}

type NamespaceResourceContainer struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []interface{} `json:"items"`
	Namespaces      []string      `json:"namespaces"`
}

// 地址拼接访问多命名空间资源
func RequestNamespaceRes(path string, ns string) string {
	ss := strings.Split(path, "/")
	resourceName := ss[len(ss)-1]
	per := ss[:len(ss)-1]
	namespacedSs := append(per, "namespaces", ns, resourceName)
	return strings.Join(namespacedSs, "/")
}

// 异步获取命名空间资源
func FetchMultiNamespaceResource(client *http.Client, namespaces []string, apiUrl url.URL) (*NamespaceResourceContainer, error) {
	wg := &sync.WaitGroup{}
	var mergedContainer NamespaceResourceContainer
	var responses []*http.Response
	var es []error
	for i := range namespaces {
		wg.Add(1)
		ns := namespaces[i]
		go func() {
			newUrl := apiUrl
			newUrl.Path = RequestNamespaceRes(apiUrl.Path, ns)
			resp, err := client.Get(newUrl.String())
			if err != nil {
				es = append(es, err)
				wg.Done()
				return
			}
			responses = append(responses, resp)
			wg.Done()
		}()

	}
	wg.Wait()
	var forbidden int
	var forbiddenMessage []string
	for i := range responses {
		r := responses[i]
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}
		if r.StatusCode != http.StatusOK {
			if r.StatusCode == http.StatusForbidden {
				forbidden++
				forbiddenMessage = append(forbiddenMessage, string(body))
				continue
			} else {
				return nil, errors.New(string(body))
			}
		}
		var nc NamespaceResourceContainer
		if err := json.Unmarshal(body, &nc); err != nil {
			return nil, err
		}
		mergedContainer.TypeMeta = nc.TypeMeta
		mergedContainer.ListMeta = nc.ListMeta
		mergedContainer.Items = append(mergedContainer.Items, nc.Items...)
	}
	if len(namespaces) == 1 && forbidden == 1 {
		return nil, errors.New(strings.Join(forbiddenMessage, ""))
	}
	return &mergedContainer, nil
}
