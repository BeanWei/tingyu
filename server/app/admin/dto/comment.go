//lint:file-ignore SA5008 .
package dto

type ListCommentReq struct {
	*Paging
	Search string `query:"search"`
	Filter struct {
		Status int8 `json:"status"`
	} `query:"filter"`
}

type UpdateCommentReq struct {
	Id     int64 `json:"id,required"`
	Status int8  `json:"status,required"`
}
