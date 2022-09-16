//lint:file-ignore SA5008 .
package types

type ListCommentReplyReq struct {
	*Paging
	Search string `query:"search"`
	Filter struct {
		Status int8 `json:"status"`
	} `query:"filter"`
}

type UpdateCommentReplyReq struct {
	Id     int64 `json:"id,required"`
	Status int8  `json:"status,required"`
}
