import  request from '../plugin/utils/request'

const listNodesUrl = (clusterName) => {
    return `/api/v1/kubernetes/${clusterName}/nodes`
}

export function ListNodes(clusterName, params) {
    let url = listNodesUrl(clusterName, params)

    return request("get", url, params)
}
