import request from "../../plugin/utils/request"

const BaseUrl = (cluster_id) => {
  return `http://localhost:19999/api/v1/kubernetes/proxy/${cluster_id}/api/v1/services`
}

const BaseNamespaceUrl = (cluster_id, namespace) => {
  return `/api/v1/kubernetes/proxy/${cluster_id}/api/v1/namespaces/${namespace}/services`
}

export const ServicesList = (cluster_id, page, pageSize, namespace, keywords, labelSelector) => {
  if (namespace && namespace !== 'All Namespaces') {
    return request(
       'get',
      `${BaseNamespaceUrl(cluster_id, namespace)}?search=true&keywords=${keywords}&page=${page}&pageSize=${pageSize}&labelSelector=${labelSelector}`,
    )
  }
  return request(
      'get',
`${BaseUrl(cluster_id)}?search=true&keywords=${keywords}&page=${page}&pageSize=${pageSize}`,)
}

export const ServicesGet = (cluster_id, namespace, name) => {
  return request(
      'get',
      `${BaseNamespaceUrl(cluster_id, namespace)}/${name}`,
      )
}

export const ServicesCreate = (cluster_id, namespace, data) => {
  return request(
      'post',
      `${BaseNamespaceUrl(cluster_id, namespace)}`,
    data)
}

export const ServicesUpdate = (cluster_id, namespace, name,  data) => {
  return request(
      'put',
 `${BaseNamespaceUrl(cluster_id, namespace)}/${name}`,
    data
  )
}

export const ServicesDelete = (cluster_id, namespace, name) => {
  return request(
      'delete',
      `${BaseNamespaceUrl(cluster_id, namespace)}/${name}`,
      )
}
