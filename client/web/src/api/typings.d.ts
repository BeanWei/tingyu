declare namespace API {
  type UserLoginReq = {
    username: string
    password: string
  }
  type UserLoginResp = {
    token: string
  }
  type GetUserInfoReq = {
    id?: string
  }
}
