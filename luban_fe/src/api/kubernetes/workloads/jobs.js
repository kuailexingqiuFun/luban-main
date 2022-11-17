import request from "../../../plugin/utils/request"

const BaseUrl = (cluster_id) => {
  return `http://localhost:19999/api/v1/kubernetes/proxy/${cluster_id}/apis/batch/v1/jobs`
}

const BaseNamespaceUrl = (cluster_id, namespace) => {
  return `/kubernetes/proxy/${cluster_id}/apis/batch/v1/namespaces/${namespace}/jobs`
}


export const JobsList = (cluster_id, page, pageSize, namespace, keywords) => {
  if (namespace && namespace !== 'All Namespaces') {
    return request(
        'get',
        `${BaseNamespaceUrl(cluster_id, namespace)}?search=true&keywords=${keywords}&page=${page}&pageSize=${pageSize}`
    )
  }
  return request(
      'get',
      `${BaseUrl(cluster_id)}?search=true&keywords=${keywords}&page=${page}&pageSize=${pageSize}`
  )
}

export const JobsGet = (cluster_id, namespace, name) => {
  return request(
      'get',
      `${BaseNamespaceUrl(cluster_id, namespace)}/${name}`
  )
}

export const JobsCreate = (cluster_id, namespace, data) => {
  return request(
      'post',
      `${BaseNamespaceUrl(cluster_id, namespace)}`,
      data)
}

export const JobsUpdate = (cluster_id, namespace, name,  data) => {
  return request(
      'put',
      `${BaseNamespaceUrl(cluster_id, namespace)}/${name}`,
      data
  )
}

export const JobsDelete = (cluster_id, namespace, name) => {
  return request(
      'delete',
      `${BaseNamespaceUrl(cluster_id, namespace)}/${name}`
  )
}
