//lint:file-ignore SA5008 .
package types

type ListTopicReq struct {
	*Paging
	SortType   int8  `query:"sort_type"`
	IsRec      bool  `query:"is_rec"`
	CategoryId int64 `query:"category_id"`
	UserId     int64 `query:"user_id"`
}

type SearchTopicReq struct {
	Keyword string `query:"keyword"`
}

type CreateTopicReq struct {
	Title       string `json:"title,required"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
}

type FollowTopicReq struct {
	Id int64 `query:"id,required"`
}

type UnFollowTopicReq struct {
	Id int64 `query:"id,required"`
}
