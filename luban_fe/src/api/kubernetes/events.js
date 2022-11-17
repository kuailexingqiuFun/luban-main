import request from "../../plugin/utils/request"

const BaseUrl = (cluster_id, namespaces) => {
  return `/api/v1/kubernetes/proxy/${cluster_id}/api/v1/namespaces/${namespaces}/events`
}

export const EventsList = (cluster_id, fieldSelector, namespaces) => {
  return request(
      'get',
      `${BaseUrl(cluster_id, namespaces)}?search=true&fieldSelector=involvedObject.name=${fieldSelector}`
  )
}
