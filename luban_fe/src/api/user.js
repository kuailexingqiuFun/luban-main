import request from "../plugin/utils/request"

export const UserLogin = (params) => request('post','http://localhost:19999/api/v1/user/login', params)
export const UserRegister = (params) => request('post', 'http://localhost:19999/api/v1/user/register', params)
export const UserManage = (params) => request('get', 'http://localhost:19999/api/v1/user/manage', params)