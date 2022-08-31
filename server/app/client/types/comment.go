//lint:file-ignore SA5008 .
package types

type ListCommentReq struct {
	*Paging
	PostId   int64 `query:"post_id,required"`
	SortType int8  `query:"sort_type"`
}

type CreateCommentReq struct {
	Content     string `json:"content,required"`
	ContentText string `json:"content_text,required"`
	PostId      int64  `json:"post_id,string,required"`
}

type DeleteCommentReq struct {
	Id int64 `query:"id,required"`
}
