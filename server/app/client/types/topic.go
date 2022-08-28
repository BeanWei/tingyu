//lint:file-ignore SA5008 .
package types

type ListTopicReq struct {
	*Paging
	SortType int8 `query:"sort_type"`
}

type CreateTopicReq struct {
	Title       string `json:"title,required"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
	Notice      string `json:"notice"`
}
