import request from "../../plugin/utils/request"

const BaseUrl = (cluster_id) => {
  return `/api/v1/kubernetes/proxy/${cluster_id}/api/v1/namespaces`
}

export const NamespaceList = (cluster_id, page, pageSize, keywords) => {
  return  request(
      "get",
      `${BaseUrl(cluster_id)}?search=true&keywords=${keywords}&page=${page}&pageSize=${pageSize}`,
      )
}

export const NamespaceCreate = (cluster_id, data) => {
  return request({
    url: `${BaseUrl(cluster_id)}`,
    method: 'post',
    data
  })
}

export const NamespaceUpdate = (cluster_id, name, data) => {
  return request({
    url: `${BaseUrl(cluster_id)}/${name}`,
    method: 'put',
    data
  })
}

export const NamespaceDelete = (cluster_id, name) => {
  return request({
    url: `${BaseUrl(cluster_id)}/${name}`,
    method: 'delete'
  })
}
