//lint:file-ignore SA5008 .
package types

type ListPostReq struct {
	*Paging
	Search string `query:"search"`
	Sorter struct {
		CommentCount int8 `json:"comment_count"`
	} `query:"sorter"`
	Filter struct {
		Status int8 `json:"status"`
	} `query:"filter"`
}

type UpdatePostReq struct {
	Id     int64 `json:"id,required"`
	Status int8  `json:"status,required"`
}
