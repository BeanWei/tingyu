//lint:file-ignore SA5008 .
package dto

type ListUserReq struct {
	*Paging
	Search string `query:"search"`
	Filter struct {
		Status  int8  `json:"status"`
		IsAdmin *int8 `json:"is_admin"`
	} `query:"filter"`
}

type UpdateUserReq struct {
	Id     int64 `json:"id,required"`
	Status int8  `json:"status,required"`
}
