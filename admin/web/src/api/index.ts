import { Message } from '@arco-design/web-react'
import axios from 'axios'
import { useUserStore } from '~/store'

export const request = axios.create({
  baseURL: '/api/admin',
  timeout: 30 * 1000, // 30s 超时
})

request.interceptors.request.use(
  (config) => {
    if (config.url && !config.url.startsWith('/')) {
      const [method, ...uri] = config.url.split(':')
      config.method = method
      config.url = uri.join(':')
    }
    !config.headers && (config.headers = { Authorization: '' })
    const token = useUserStore.getState().token
    if (token)
      config.headers.Authorization = `Bearer ${token}`
    return config
  },
  (error) => {
    return Promise.reject(error)
  },
)

request.interceptors.response.use(
  (response) => {
    if (response.status === 200)
      return response
    else
      Promise.reject(response?.data || {})
  },
  (error = {}) => {
    const { response = {} } = error || {}
    // 重定向
    if (response?.status === 401)
      useUserStore.getState().deleteToken()
    else if (response?.status === 403)
      Message.warning(response?.data.msg || '禁止访问')
    else
      Message.error(response?.data?.msg || '请求失败')
    return Promise.reject(response?.data || {})
  },
)

export interface Result<T> {
  code: number
  msg: string
  total: number
  data: T
}

export const url = {
  userLogin: 'POST:/auth/login',
  getUserInfo: 'GET:/auth/info',
  listCategory: 'GET:/category/list',
  createCategory: 'POST:/category/create',
  updateCategory: 'PUT:/category/update',
  listTopic: 'GET:/topic/list',
  createTopic: 'POST:/topic/create',
  updateTopic: 'PUT:/topic/update',
  listPost: 'GET:/post/list',
  updatePost: 'PUT:/post/update',
  listComment: 'GET:/comment/list',
  updateComment: 'PUT:/comment/update',
  listReply: 'GET:/reply/list',
  updateReply: 'PUT:/reply/update',
  listUser: 'GET:/user/list',
  updateUser: 'PUT:/user/update',
}
