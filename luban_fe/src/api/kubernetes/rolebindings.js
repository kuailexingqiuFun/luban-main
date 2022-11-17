import request from "../../plugin/utils/request"


const BaseUrl = (cluster_id) => {
  return `http://localhost:19999/api/v1/kubernetes/proxy/${cluster_id}/apis/rbac.authorization.k8s.io/v1/rolebindings`
}

const BaseNamespaceUrl = (cluster_id, namespace) => {
  return `http://localhost:19999/api/v1/kubernetes/proxy/${cluster_id}/apis/rbac.authorization.k8s.io/v1/namespaces/${namespace}/rolebindings`
}

export const RoleBindingsList = (cluster_id, page, pageSize, namespace, keywords, labelSelector) => {
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

export const RoleBindingsGet = (cluster_id, namespace, name) => {
  return request(
      'get',
      `${BaseNamespaceUrl(cluster_id, namespace)}/${name}`,
      )
}

export const RoleBindingsCreate = (cluster_id, namespace, data) => {
  return request(
      'post',
      `${BaseNamespaceUrl(cluster_id, namespace)}`,
      data
  )
}

export const RoleBindingsUpdate = (cluster_id, namespace, name,  data) => {
  return request(
      'put',
      `${BaseNamespaceUrl(cluster_id, namespace)}/${name}`,
      data
  )
}

export const RoleBindingsDelete = (cluster_id, namespace, name) => {
  return request(
      'delete',
      `${BaseNamespaceUrl(cluster_id, namespace)}/${name}`,
      )
}
