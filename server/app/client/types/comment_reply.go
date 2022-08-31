//lint:file-ignore SA5008 .
package types

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

type DeleteCommentReplyReq struct {
	Id int64 `query:"id,required"`
}
