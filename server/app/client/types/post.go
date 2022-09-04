//lint:file-ignore SA5008 .
package types

type ListPostReq struct {
	*Paging
	SortType int8  `query:"sort_type"`
	TopicId  int64 `query:"topic_id"`
}

type GetPostReq struct {
	Id int64 `query:"id,required"`
}

type CreatePostReq struct {
	Content     string  `json:"content,required"`
	ContentText string  `json:"content_text,required"`
	TopicIds    []int64 `json:"topic_ids"`
}

type SearchPostReq struct {
	*Paging
	Keyword string `query:"keyword,required"`
}
