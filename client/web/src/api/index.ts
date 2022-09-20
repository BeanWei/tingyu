import type { UploadCustomRequestOptions } from 'naive-ui'
import axios from 'axios'

export const instance = axios.create({
  baseURL: '/api',
  timeout: 30 * 1000, // 30s 超时
})

instance.interceptors.request.use(
  (config) => {
    if (config.url && !config.url.startsWith('/')) {
      const [method, ...uri] = config.url.split(':')
      config.method = method
      config.url = uri.join(':')
    }
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

export const url = {
  userLogin: 'POST:/v1/user/login',
  getUserInfo: 'GET:/v1/user/get',
  updateUserInfo: 'PUT:/v1/user/update',
  fileUpload: 'POST:/v1/upload',
  listCategory: 'GET:/v1/category/list',
  listTopic: 'GET:/v1/topic/list',
  searchTopic: 'GET:/v1/topic/search',
  createTopic: 'POST:/v1/topic/create',
  followTopic: 'POST:/v1/topic/follow',
  unFollowTopic: 'DELETE:/v1/topic/unfollow',
  listPost: 'GET:/v1/post/list',
  searchPost: 'GET:/v1/post/search',
  getPost: 'GET:/v1/post/get',
  createPost: 'POST:/v1/post/create',
  listComment: 'GET:/v1/comment/list',
  createComment: 'POST:/v1/comment/create',
  deleteComment: 'DELETE:/v1/comment/delete',
  listReply: 'GET:/v1/reply/list',
  createReply: 'POST:/v1/reply/create',
  deleteReply: 'DELETE:/v1/reply/delete',
}

export const fileUpload = (options: UploadCustomRequestOptions) => {
  const {
    file,
    data,
    onFinish,
    onError,
    onProgress,
  } = options
  const formData = new FormData()
  if (data) {
    Object.keys(data).forEach((key) => {
      formData.append(
        key,
        data[key as keyof UploadCustomRequestOptions['data']],
      )
    })
  }
  formData.append('file', file.file as File)
  return instance({
    url: `${url.fileUpload}`,
    data: formData,
    onUploadProgress: ({ percent }) => {
      onProgress({ percent: Math.ceil(percent) })
    },
  }).then((result) => {
    file.url = result.data.data.url
    file.thumbnailUrl = result.data.data.url
    onFinish()
  }).catch(() => {
    onError()
  })
}
