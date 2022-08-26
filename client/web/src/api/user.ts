import type { Promise } from './base'
import { api } from './base'

export const userLogin = (data: API.UserLoginReq): Promise<API.UserLoginResp> => {
  return api.execute('/v1/user/login', {
    method: 'post',
    data,
  })
}

export const getUserInfo = (params: API.GetUserInfoReq) => {
  return api.execute('/v1/user/get', {
    method: 'get',
    params,
  })
}
