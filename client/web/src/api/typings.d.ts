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
}
