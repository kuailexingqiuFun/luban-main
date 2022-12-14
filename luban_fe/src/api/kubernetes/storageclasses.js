import request from "../../plugin/utils/request"

const BaseUrl = (cluster_id) => {
  return `/api/v1/kubernetes/proxy/${cluster_id}/apis/storage.k8s.io/v1/storageclasses`
}

const BaseNamespaceUrl = (cluster_id) => {
  return `/api/v1/kubernetes/proxy/${cluster_id}/apis/storage.k8s.io/v1/storageclasses`
}

export const StorageClassList = (cluster_id, page, pageSize, namespace, keywords, labelSelector) => {
  if (namespace && namespace !== 'All Namespaces') {
    return request(
        'get',
        `${BaseNamespaceUrl(cluster_id, namespace)}?search=true&keywords=${keywords}&page=${page}&pageSize=${pageSize}&labelSelector=${labelSelector}`,
        )
  }
  return request(
      'get',
      `${BaseUrl(cluster_id)}?search=true&keywords=${keywords}&page=${page}&pageSize=${pageSize}`,
      )
}

export const StorageClassGet = (cluster_id, namespace, name) => {
  return request(
      'get',
      `${BaseNamespaceUrl(cluster_id, namespace)}/${name}`)
}

export const StorageClassCreate = (cluster_id, namespace, data) => {
  return request(
      'post',
      `${BaseNamespaceUrl(cluster_id, namespace)}`,
      data
  )
}

export const StorageClassUpdate = (cluster_id, namespace, name, data) => {
  return request(
      'put',
      `${BaseNamespaceUrl(cluster_id, namespace)}/${name}`,
      data
  )
}

export const StorageClassDelete = (cluster_id, namespace, name) => {
  return request(
      'delete',
      `${BaseNamespaceUrl(cluster_id, namespace)}/${name}`
  )
}
