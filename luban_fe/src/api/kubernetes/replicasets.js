import request from "../../plugin/utils/request"

const BaseUrl = (cluster_id, namespace) => {
  return `/api/v1/kubernetes/proxy/${cluster_id}/apis/apps/v1/namespaces/${namespace}/replicasets`
}

export function ReplicaSetsList(cluster_id, namespace, selector) {
  return request(
      'get',
    `${BaseUrl(cluster_id, namespace)}?search=true&labelSelector=${selector}+&namespace=${namespace}`,
      )
}
