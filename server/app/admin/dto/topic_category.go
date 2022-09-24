//lint:file-ignore SA5008 .
package dto

type ListTopicCategoryReq struct {
	*Paging
	Search string `query:"search"`
	Sorter struct {
		Rank int8 `json:"rank"`
	} `query:"sorter"`
}

type CreateTopicCategoryReq struct {
	Name string `json:"name,required"`
	Rank int    `json:"rank"`
}

type UpdateTopicCategoryReq struct {
	Id   int64  `json:"id,required"`
	Name string `json:"name,required"`
	Rank int    `json:"rank"`
}
