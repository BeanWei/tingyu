//lint:file-ignore SA5008 .
package types

type ListTopicReq struct {
	*Paging
	Search string `query:"search"`
	Sorter struct {
		PostCount     int8 `json:"post_count"`
		FollowerCount int8 `json:"follower_count"`
		AttenderCount int8 `json:"attender_count"`
	} `query:"sorter"`
	Filter struct {
		Status int8 `json:"status"`
	} `query:"filter"`
}

type CreateTopicReq struct {
	Title           string `json:"title,required"`
	Icon            string `json:"icon"`
	Description     string `json:"description"`
	Status          int8   `json:"status"`
	TopicCategoryId int64  `json:"topic_category_id,required"`
	IsRec           bool   `json:"is_rec"`
	RecRank         int    `json:"rec_rank"`
}

type UpdateTopicReq struct {
	Id              int64  `json:"id,required"`
	Title           string `json:"title,required"`
	Icon            string `json:"icon"`
	Description     string `json:"description"`
	Status          int8   `json:"status"`
	TopicCategoryId int64  `json:"topic_category_id,required"`
	IsRec           bool   `json:"is_rec"`
	RecRank         int    `json:"rec_rank"`
}
