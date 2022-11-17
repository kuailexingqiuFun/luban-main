import request from "../../plugin/utils/request"

const BaseUrl = (cluster_id) => {
  return `/api/v1/kubernetes/proxy/${cluster_id}/api/v1/persistentvolumes`
}

const BaseNamespaceUrl = (cluster_id, namespace) => {
  return `/api/v1/kubernetes/proxy/${cluster_id}/api/v1/persistentvolumes`
}

export const PersistentVolumesList = (cluster_id, page, pageSize, namespace, keywords, labelSelector) => {
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

export const PersistentVolumesGet = (cluster_id, namespace, name) => {
  return request(
      'get',
      `${BaseNamespaceUrl(cluster_id, namespace)}/${name}`
  )
}

export const PersistentVolumesCreate = (cluster_id, namespace, data) => {
  return request(
      'post',
      `${BaseNamespaceUrl(cluster_id, namespace)}`,
      data
  )
}

export const PersistentVolumesUpdate = (cluster_id, namespace, name, data) => {
  return request(
      'put',
      `${BaseNamespaceUrl(cluster_id, namespace)}/${name}`,
      data
  )
}

export const PersistentVolumesDelete = (cluster_id, namespace, name) => {
  return request(
      'delete',
      `${BaseNamespaceUrl(cluster_id, namespace)}/${name}`
  )
}