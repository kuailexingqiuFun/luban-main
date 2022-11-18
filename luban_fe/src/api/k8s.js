import  request from '../plugin/utils/request'

const listNodesUrl = (clusterName) => {
    return `/api/v1/kubernetes/${clusterName}/nodes`
}

export function ListNodes(clusterName, params) {
    let url = listNodesUrl(clusterName, params)
    return request("get", url, params)
}

export function GetNodeDetail(clusterName, nodeName) {
    let url = listNodesUrl(clusterName) + "/" + nodeName
    return request("get", url)
}

export function DeleteCluster(params) {
    return request('delete', '/api/v1/kubernetes/cluster', params)
}

export function EditCluster(clusterId, params) {
    return request('put', `/api/v1/kubernetes/clusters/${clusterId}`, params)
}

export function GetClusterDetail(clusterId) {
    return request('get', `/api/v1/kubernetes/clusters/${clusterId}`)
}