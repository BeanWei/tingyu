import type { P } from './base'
import { api } from './base'

export const createPost = (data: API.CreatePostReq) => {
  return api.execute('/v1/post/create', {
    method: 'post',
    data,
  })
}

export const listPost = (params?: API.ListPostReq): P<API.ListPostResp> => {
  return api.execute('/v1/post/list', {
    method: 'get',
    params,
  })
}