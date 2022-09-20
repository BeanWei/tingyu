//lint:file-ignore SA5008 .
package types

type UserLoginReq struct {
	Username string `json:"username,required"`
	Password string `json:"password,required"`
}

type GetUserInfoReq struct {
	Id int64 `query:"id"`
}

type UpdateUserInfoReq struct {
	Avatar   string `json:"avatar"`
	Nickname string `json:"nickname,required"`
	Headline string `json:"headline"`
}
