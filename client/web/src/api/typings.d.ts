declare namespace API {
  type UserLoginReq = {
    username: string
    password: string
  }
  type UserLoginResp = {
    token: string
  }
  type GetUserInfoReq = {
    id?: number
  }
  type GetUserInfoResp = {
    id: number
    username: string
    nickname: string
    avatar: string
  }
  type CreatePostReq = {
    content: string
    content_text: string
    topic_ids?: number[]
  }
  type ListPostReq = {
    limit: number
    page: number
    sort_type?: number
    topic_id?: number
  }
  type ListPostResp = {
    id: number
    created_at: number
    updated_at: number
    content: string
    comment_count: number
    is_top: boolean
    is_excellent: boolean
    is_lock: boolean
    latest_replied_at: number
    ip_loc: string
    user: object
  }
}
