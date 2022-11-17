import request from "../../plugin/utils/request"

export const getClusterList = (params) =>  {
   return  request("get", "/api/v1/kubernetes/clusters", params)
}

export const createCluster = (params) => {
   return request("post", "/api/v1/kubernetes/cluster", params)
}