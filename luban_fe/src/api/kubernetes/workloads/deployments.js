import request from "../../../plugin/utils/request"

const BaseUrl = (cluster_id) => {
  return `http://localhost:19999/api/v1/kubernetes/proxy/${cluster_id}/apis/apps/v1/deployments`
}

const BaseNamespaceUrl = (cluster_id, namespace) => {
  return `http://localhost:19999/api/v1/kubernetes/proxy/${cluster_id}/apis/apps/v1/namespaces/${namespace}/deployments`
}

export const DeploymentsList = (cluster_id, page, pageSize, namespace, keywords) => {
  if (namespace && namespace !== 'All Namespaces') {
    return  request(
        "get",
        `${BaseNamespaceUrl(cluster_id, namespace)}?search=true&keywords=${keywords}&page=${page}&pageSize=${pageSize}`,
    )
  }

  return  request(
      "get",
      `${BaseUrl(cluster_id)}?search=true&keywords=${keywords}&page=${page}&pageSize=${pageSize}`,
  )
}

export const DeploymentsGet = (cluster_id, namespace, name) => {
  return  request(
      "get",
      `${BaseNamespaceUrl(cluster_id, namespace)}/${name}`,
  )
}

export const DeploymentsCreate = (cluster_id, namespace, data) => {
  return request({
    url: `${BaseNamespaceUrl(cluster_id, namespace)}`,
    method: 'post',
    data
  })
}

export const DeploymentsUpdate = (cluster_id, namespace, name, data) => {
  return  request(
      "put",
      `${BaseNamespaceUrl(cluster_id, namespace)}/${name}?namespace=${namespace}`,
      data
  )
}

export const DeploymentsPatch = (cluster_id, namespace, name, data) => {
  return request({
    url: `${BaseNamespaceUrl(cluster_id, namespace)}/${name}?namespace=${namespace}`,
    method: 'patch',
    data
  })
}

export const DeploymentsDelete = (cluster_id, namespace, name) => {
  return request({
    url: `${BaseNamespaceUrl(cluster_id, namespace)}/${name}`,
    method: 'delete'
  })
}