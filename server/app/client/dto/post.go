//lint:file-ignore SA5008 .
package dto

import (
	"github.com/BeanWei/tingyu/data/ent"
	"github.com/BeanWei/tingyu/data/ent/userreaction"
)

type Post struct {
	ID              int64       `json:"id"`
	CreatedAt       int64       `json:"created_at"`
	UpdatedAt       int64       `json:"updated_at"`
	CommentCount    int         `json:"comment_count"`
	IsTop           bool        `json:"is_top"`
	IsExcellent     bool        `json:"is_excellent"`
	IsLock          bool        `json:"is_lock"`
	LatestRepliedAt int64       `json:"latest_replied_at"`
	IPLoc           string      `json:"ip_loc"`
	Content         string      `json:"content"`
	UserID          int64       `json:"user_id"`
	User            *ent.User   `json:"user,omitempty"`
	Reactions       []*Reaction `json:"reactions"`
}

type ListPostReq struct {
	*Paging
	SortType int8  `query:"sort_type"`
	TopicId  int64 `query:"topic_id"`
	UserId   int64 `query:"user_id"`
	Pinned   bool  `query:"pinned"`
	Starred  bool  `query:"starred"`
}

type SearchPostReq struct {
	*Paging
	Keyword string `query:"keyword,required"`
}

type GetPostReq struct {
	Id int64 `query:"id,required"`
}

type CreatePostReq struct {
	Content     string  `json:"content,required"`
	ContentText string  `json:"content_text,required"`
	TopicIds    []int64 `json:"topic_ids"`
}

type ReactPostReq struct {
	Id   int64                  `json:"id,required"`
	Code userreaction.ReactCode `json:"code,required"`
}
