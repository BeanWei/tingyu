//lint:file-ignore SA5008 .
package dto

import (
	"github.com/BeanWei/tingyu/data/ent"
	"github.com/BeanWei/tingyu/data/ent/userreaction"
)

type Comment struct {
	ID             int64           `json:"id"`
	CreatedAt      int64           `json:"created_at"`
	UpdatedAt      int64           `json:"updated_at"`
	PostID         int64           `json:"post_id"`
	UserID         int64           `json:"user_id"`
	IPLoc          string          `json:"ip_loc"`
	Content        string          `json:"content"`
	ReplyCount     int             `json:"reply_count"`
	IsPoster       bool            `json:"is_poster"`
	User           *ent.User       `json:"user"`
	Reactions      []*Reaction     `json:"reactions"`
	CommentReplies []*CommentReply `json:"comment_replies"`
}

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

type ReactCommentReq struct {
	Id   int64                  `json:"id,required"`
	Code userreaction.ReactCode `json:"code,required"`
}

type DeleteCommentReq struct {
	Id int64 `query:"id,required"`
}
