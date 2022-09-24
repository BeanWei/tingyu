//lint:file-ignore SA5008 .
package dto

type UserLoginReq struct {
	Username string `json:"username,required"`
	Password string `json:"password,required"`
}
