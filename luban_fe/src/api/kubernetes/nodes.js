import request from "../../plugin/utils/request"

const BaseUrl = (cluster_id) => {
  return `/api/v1/kubernetes/proxy/${cluster_id}/api/v1/nodes`
}

export const NodesList = (cluster_id, page, pageSize, keywords) => {
    return request(
        'get',
        `${BaseUrl(cluster_id)}?search=true&keywords=${keywords}&page=${page}&pageSize=${pageSize}`,
        )
}

export const NodesGet = (cluster_id, name) => {
  return request(
      'get',
      `${BaseUrl(cluster_id)}/${name}`,
      )
}

export const NodesUpdate = (cluster_id, name, data) => {
  return request(
      'put',
      `${BaseUrl(cluster_id)}/${name}`,
      data
  )
}

export const NodesDelete = (cluster_id, name) => {
  return request(
      'delete',
      `${BaseUrl(cluster_id)}/${name}`
  )
}
