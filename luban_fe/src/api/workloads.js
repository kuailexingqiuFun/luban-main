import request from "../plugin/utils/request"

const BaseUrl = () => {
    return `/api/v1/kubernetes/proxy/`
}
// const apiV1Url = (cluster_id, type) => {
//     return `/k8s/proxy/${cluster_id}/api/v1/${type}`
// }
const apiV1UrlWithNsUrl = (cluster_id, type, namespaces) => {
    return `${BaseUrl()}${cluster_id}/api/v1/namespaces/${namespaces}/${type}`
}
// const appsV1Url = (cluster_id, type) => {
//     return `/api/v1/proxy/${cluster_id}/k8s/apis/apps/v1/${type}`
// }
const appsV1UrlWithNsUrl = (cluster_id, type, namespaces) => {
    return `${BaseUrl()}${cluster_id}/apis/apps/v1/namespaces/${namespaces}/${type}`
}
// const batchV1beta1Url = (cluster_id, type) => {
//     return `/k8s/proxy/${cluster_id}/apis/batch/v1beta1/${type}`
// }
const batchV1beta1WithNsUrl = (cluster_id, type, namespaces) => {
    return `${BaseUrl}${cluster_id}/apis/batch/v1beta1/namespaces/${namespaces}/${type}`
}
// const batchV1Url = (cluster_id, type) => {
//     return `/k8s/proxy/${cluster_id}/apis/batch/v1/${type}`
// }

const batchV1WithNsUrl = (cluster_id, type, namespaces) => {
    return `${BaseUrl()}${cluster_id}/apis/batch/v1/namespaces/${namespaces}/${type}`
}

const apiNsUrl = (cluster_id, version, namespace, kind) => {
    return `$${BaseUrl()}${cluster_id}/api/${version}/namespaces/${namespace}/${kind}`
}

const apiUrl = (cluster_id, version, kind) => {
    return `${BaseUrl()}${cluster_id}/api/${version}/${kind}`
}

const apisNsUrl = (cluster_id, version, group, namespace, kind) => {
    return `${BaseUrl()}${cluster_id}/apis/${group}/${version}/namespaces/${namespace}/${kind}`
}

const apisUrl = (cluster_id, version, group, kind) => {
    return `${BaseUrl()}${cluster_id}/apis/${group}/${version}/${kind}`
}

export function postYaml (cluster_id, kind, data) {
    return request(
        'post',
        `${getUrl(cluster_id, kind, data)}`
  )
}

export function getUrl (cluster_id, kind, data) {
    let url = ""
    if (data.apiVersion === "v1") {
        if (data.metadata.namespace !== undefined && data.metadata.namespace !== "") {
            url = apiNsUrl(cluster_id, data.apiVersion, data.metadata.namespace, kind)
        } else {
            url = apiUrl(cluster_id, data.apiVersion, kind)
        }
    } else {
        const apiVersions = data.apiVersion.split("/")
        const group = apiVersions[0]
        const version = apiVersions[1]
        if (data.metadata.namespace !== undefined && data.metadata.namespace !== "") {
            url = apisNsUrl(cluster_id, version, group, data.metadata.namespace, kind)
        } else {
            url = apisUrl(cluster_id, version, group, kind)
        }
    }
    return url
}

export function getWorkLoadByName (cluster_id, type, namespace, name) {
    switch (type) {
        case "deployments":
        case "statefulsets":
        case "daemonsets":
            return request(
                'get',
                `${appsV1UrlWithNsUrl(cluster_id, type, namespace)}/${name}`
            )
        case "cronjobs":
            return request(
                    'get',
                `${batchV1beta1WithNsUrl(cluster_id, type, namespace)}/${name}`
            )
        case "pods":
            return request(
                    'get',
                `${apiV1UrlWithNsUrl(cluster_id, type, namespace)}/${name}`
            )
        case "jobs":
            return request(
                'get',
                `${batchV1WithNsUrl(cluster_id, type, namespace)}/${name}`
            )
    }
}

export function createWorkLoad (cluster_id, type, namespace, data) {
    switch (type) {
        case "deployments":
        case "statefulsets":
        case "daemonsets":
            return request(
                 'post',
                 `${appsV1UrlWithNsUrl(cluster_id, type, namespace)}`,
                data
            )
        case "cronjobs":
            return request(
                'post',
            `${batchV1beta1WithNsUrl(cluster_id, type, namespace)}`,
                data)
        case "pods":
        case "services":
            return request(
                'post',
            `${apiV1UrlWithNsUrl(cluster_id, type, namespace)}`,
                data
            )
        case "jobs":
            return request(
                'post',
            `${batchV1WithNsUrl(cluster_id, type, namespace)}`,
                data
            )
    }
}

export function deleteWorkLoad (cluster_id, type, namespace, name) {
    switch (type) {
        case "deployments":
        case "statefulsets":
        case "daemonsets":
            return request(
                'delete',
            `${appsV1UrlWithNsUrl(cluster_id, type, namespace)}/${name}`
            )
        case "cronjobs":
            return service({
                url: `${batchV1beta1WithNsUrl(cluster_id, type, namespace)}/${name}`,
                method: 'delete',
            })
        case "pods":
            return service({
                url: `${apiV1UrlWithNsUrl(cluster_id, type, namespace)}/${name}`,
                method: 'delete',
            })
        case "jobs":
            return service({
                url: `${batchV1WithNsUrl(cluster_id, type, namespace)}/${name}`,
                method: 'delete',
            })
    }
}

export function updateWorkLoad (cluster_id, type, namespace, name, data) {
    switch (type) {
        case "deployments":
        case "statefulsets":
        case "daemonsets":

            return request(
                'put', `${appsV1UrlWithNsUrl(cluster_id, type, namespace)}/${name}`,
                data
            )
        case "cronjobs":
            return request(
                'put',
                `${batchV1beta1WithNsUrl(cluster_id, type, namespace)}/${name}`,
                data
               )
        case "pods":
            return request(
                'put',
                `${apiV1UrlWithNsUrl(cluster_id, type, namespace)}/${name}`,
                data
            )
        case "jobs":
            return request(
                'put',
                `${batchV1WithNsUrl(cluster_id, type, namespace)}/${name}`,
                data
            )
    }
}
