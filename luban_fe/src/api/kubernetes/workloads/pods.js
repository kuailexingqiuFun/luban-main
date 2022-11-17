import request from "../../../plugin/utils/request"

const BaseUrl = (cluster_id) => {
  return `/api/v1/kubernetes/proxy/${cluster_id}/api/v1/pods`
}

const BaseNamespaceUrl = (cluster_id, namespace) => {
  return `/api/v1/kubernetes/proxy/${cluster_id}/api/v1/namespaces/${namespace}/pods`
}

const EvictionUrl = (cluster_id, namespace, name) => {
  return `/api/v1/kubernetes/proxy/${cluster_id}/api/v1/namespaces/${namespace}/pods/${name}/eviction`
}

export const PodsList = (cluster_id, page, pageSize, namespace, keywords, labelSelector, fieldSelector) => {
  if (namespace && namespace !== 'All Namespaces') {
    return request(
        'get',
      `${BaseNamespaceUrl(cluster_id, namespace)}?search=true&keywords=${keywords}&page=${page}&pageSize=${pageSize}&labelSelector=${labelSelector}&fieldSelector=${fieldSelector}`,
    )
  }
  return request(
    'get',
    `${BaseUrl(cluster_id)}?search=true&keywords=${keywords}&page=${page}&pageSize=${pageSize}&labelSelector=${labelSelector}&fieldSelector=${fieldSelector}`,
  )
}

export const PodsGet = (cluster_id, namespace, name) => {
  return request(
    'get',
`${BaseNamespaceUrl(cluster_id, namespace)}/${name}`
  )
}

export const PodsCreate = (cluster_id, namespace, data) => {
  return request(
    'post',
    `${BaseNamespaceUrl(cluster_id, namespace)}`,
      data
  )
}

export const PodsUpdate = (cluster_id, namespace, name, data) => {
  return request(
    'put',
    `${BaseNamespaceUrl(cluster_id, namespace)}/${name}`,
    data
  )
}

export const PodsDelete = (cluster_id, namespace, name) => {
  return request(
      'delete',
    `${BaseNamespaceUrl(cluster_id, namespace)}/${name}`,
   )
}

export const Podseviction = (cluster_id, name, namespaces, data) => {
  return request(
      'post',
    `${EvictionUrl(cluster_id, namespaces, name)}`,
    data)
}
