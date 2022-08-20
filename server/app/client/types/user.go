//lint:file-ignore SA5008 .
package types

type UserLoginReq struct {
	Username string `json:"username,required"`
	Password string `json:"password,required"`
}

type UserSignInReq struct {
	Username string `json:"username,required"`
	Password string `json:"password,required"`
}

type GetUserInfoReq struct {
	ID int64 `query:"id"`
}
