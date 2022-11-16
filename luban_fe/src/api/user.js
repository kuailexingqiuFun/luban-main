import request from "../plugin/utils/request"

export const UserLogin = (params) => request('post','http://localhost:19999/api/v1/user/login', params)