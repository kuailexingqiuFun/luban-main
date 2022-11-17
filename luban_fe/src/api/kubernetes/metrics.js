import request from "../../plugin/utils/request"

const BaseUrl = (cluster_id) => {
  return `http://localhost:19999/api/v1/kubernetes/proxy/${cluster_id}`
}

export const NodeMetricsList = (cluster_id, data) => {
  return request(
      'get',
    `${BaseUrl(cluster_id)}/apis/metrics.k8s.io/v1beta1/nodes?search=true`,
    data)
}

export const PodsMetricsList = (cluster_id, namespace, labelSelector) => {
  return request(
      'get',
    `${BaseUrl(cluster_id)}/apis/metrics.k8s.io/v1beta1/pods?search=true&fieldSelector=metadata.namespace=${namespace}&labelSelector=${labelSelector}&namespace=${namespace}`,
  )
}
