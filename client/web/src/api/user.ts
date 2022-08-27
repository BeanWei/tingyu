import type { P } from './base'
import { api } from './base'

export const userLogin = (data: API.UserLoginReq): P<API.UserLoginResp> => {
  return api.execute('/v1/user/login', {
    method: 'post',
    data,
  })
}

export const getUserInfo = (params?: API.GetUserInfoReq): P<API.GetUserInfoResp> => {
  return api.execute('/v1/user/get', {
    method: 'get',
    params,
  })
}
