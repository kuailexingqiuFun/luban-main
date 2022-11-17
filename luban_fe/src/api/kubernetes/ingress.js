import request from "../../plugin/utils/request"

const BaseUrl = (cluster_id) => {
  return `http://localhost:19999/api/v1/kubernetes/proxy/${cluster_id}/apis/networking.k8s.io/v1/ingresses`
}

const BaseNamespaceUrl = (cluster_id, namespace) => {
  return `http://localhost:19999/api/v1/kubernetes/proxy/${cluster_id}/apis/networking.k8s.io/v1/namespaces/${namespace}/ingresses`
}

export const IngressesList = (cluster_id, page, pageSize, namespace, keywords, labelSelector, fieldSelector) => {
  if (namespace && namespace !== 'All Namespaces') {
    return request(
        'get',
        `${BaseNamespaceUrl(cluster_id, namespace)}?search=true&keywords=${keywords}&page=${page}&pageSize=${pageSize}`,
    )
  }
  return request(
      'get',
      `${BaseUrl(cluster_id)}?search=true&keywords=${keywords}&page=${page}&pageSize=${pageSize}`,
  )
}

export const IngressesGet = (cluster_id, namespace, name) => {
  return request(
      'get',
      `${BaseNamespaceUrl(cluster_id, namespace)}/${name}`,
      )
}

export const IngressesCreate = (cluster_id, namespace, data) => {
  return request(
      'post',
      `${BaseNamespaceUrl(cluster_id, namespace)}`,
      data
  )
}

export const IngressesUpdate = (cluster_id, namespace, name, data) => {
  return request(
      'put',
      `${BaseNamespaceUrl(cluster_id, namespace)}/${name}`,
    data
  )
}

export const IngressesDelete = (cluster_id, namespace, name) => {
  return request(
      'delete',
      `${BaseNamespaceUrl(cluster_id, namespace)}/${name}`,
      )
}
