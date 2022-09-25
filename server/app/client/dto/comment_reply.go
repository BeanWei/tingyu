//lint:file-ignore SA5008 .
package dto

import (
	"github.com/BeanWei/tingyu/data/ent"
	"github.com/BeanWei/tingyu/data/ent/userreaction"
)

type CommentReply struct {
	ID        int64       `json:"id"`
	CreatedAt int64       `json:"created_at"`
	UpdatedAt int64       `json:"updated_at"`
	UserID    int64       `json:"user_id"`
	IPLoc     string      `json:"ip_loc"`
	Content   string      `json:"content"`
	CommentID int64       `json:"comment_id"`
	ToUserID  int64       `json:"to_user_id"`
	ToReplyID int64       `json:"to_reply_id"`
	IsPoster  bool        `json:"is_poster"`
	User      *ent.User   `json:"user"`
	Reactions []*Reaction `json:"reactions"`
}

type ListCommentReplyReq struct {
	*Paging
	CommentId int64 `query:"comment_id,required"`
}

type CreateCommentReplyReq struct {
	Content     string `json:"content,required"`
	ContentText string `json:"content_text,required"`
	CommentId   int64  `json:"comment_id,required"`
	ToUserId    int64  `json:"to_user_id"`
	ToReplyId   int64  `json:"to_reply_id"`
}

type ReactCommentReplyReq struct {
	Id   int64                  `json:"id,required"`
	Code userreaction.ReactCode `json:"code,required"`
}

type DeleteCommentReplyReq struct {
	Id int64 `query:"id,required"`
}
