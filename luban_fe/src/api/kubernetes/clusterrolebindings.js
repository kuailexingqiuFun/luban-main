import request from "../../plugin/utils/request"

const BaseUrl = (cluster_id) => {
  return `/api/v1/kubernetes/proxy/${cluster_id}/apis/rbac.authorization.k8s.io/v1/clusterrolebindings`
}

const BaseNamespaceUrl = (cluster_id, namespace) => {
  return `/api/v1/kubernetes/proxy/${cluster_id}/apis/rbac.authorization.k8s.io/v1/clusterrolebindings`
}

export const ClusterRoleBindingsList = (cluster_id, page, pageSize, namespace, keywords, labelSelector) => {
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

export const ClusterRoleBindingsGet = (cluster_id, namespace, name) => {
  return request(
      'get',
      `${BaseNamespaceUrl(cluster_id, namespace)}/${name}`,
      )
}

export const ClusterRoleBindingsCreate = (cluster_id, namespace, data) => {
  return request(
      'post',
      `${BaseNamespaceUrl(cluster_id, namespace)}`,
      data
  )
}

export const ClusterRoleBindingsUpdate = (cluster_id, namespace, name,  data) => {
  return request(
      'put',
      `${BaseNamespaceUrl(cluster_id, namespace)}/${name}`,
      data
  )
}

export const ClusterRoleBindingsDelete = (cluster_id, namespace, name) => {
  return request(
      'delete',
      `${BaseNamespaceUrl(cluster_id, namespace)}/${name}`,
      )
}
