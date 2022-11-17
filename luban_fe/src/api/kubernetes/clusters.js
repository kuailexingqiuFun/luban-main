import request from "../../plugin/utils/request"

export const getClusterList = (params) =>  {
   return  request(
        "get",
        "http://localhost:19999/api/v1/kubernetes/clusters",
        params)
}

export const createCluster = (params) => {
   return request(
       "post",
       "http://localhost:19999/api/v1/kubernetes/cluster",
       params
   )
}