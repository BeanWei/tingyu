import axios from 'axios'
import type { StrictUseAxiosReturn } from '@vueuse/integrations/useAxios'
import { useAxios } from '@vueuse/integrations/useAxios'

const instance = axios.create({
  baseURL: '/api',
  timeout: 30 * 1000, // 30s 超时
})

instance.interceptors.request.use(
  (config) => {
    !config.headers && (config.headers = { Authorization: '' })
    const token = useUserStore().token
    if (token)
      config.headers.Authorization = `Bearer ${token}`
    return config
  },
  (error) => {
    return Promise.reject(error)
  },
)

instance.interceptors.response.use(
  (response) => {
    if (response.status === 200)
      return response
    else
      Promise.reject(response?.data || {})
  },
  (error = {}) => {
    const { response = {} } = error || {}
    // 重定向
    if (response?.status === 401) {
      useUserStore().resetAll()
      useMiscStore().setAuthModalVisible(true)
    }
    else if (response?.status === 403) {
      window.$message?.warning(response?.data.msg || '禁止访问')
    }
    else {
      window.$message?.error(response?.data?.msg || '请求失败')
    }

    return Promise.reject(response?.data || {})
  },
)

export interface Result<T> {
  code: number
  msg: string
  total: number
  data: T
}

export interface P<T> extends PromiseLike<StrictUseAxiosReturn<Result<T>>> {}

export const api = useAxios<Result<any>>('', instance, { immediate: false })
