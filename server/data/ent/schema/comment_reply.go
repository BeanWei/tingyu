package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/BeanWei/tingyu/pkg/entx/mixin"
)

type CommentReply struct {
	ent.Schema
}

func (CommentReply) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.ID{},
		mixin.Time{},
	}
}

func (CommentReply) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").Comment("用户ID"),
		field.String("ip").Default("").Sensitive().Comment("IP"),
		field.String("ip_loc").Default("").Comment("IP地址"),
		field.String("content").Default("").Comment("回复内容"),
		field.Int64("comment_id").Comment("评论ID"),
		field.Int64("to_user_id").Default(0).Comment("回复的用户ID"),
		field.Int64("to_reply_id").Default(0).Comment("回复的回复ID"),
		field.Bool("is_poster").Default(false).Comment("是否发帖者"),
		field.Int8("status").Default(0).Comment("状态(1.审核通过 2.审核未通过 3.待审核)"),
	}
}

func (CommentReply) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("comment_id"),
	}
}

func (CommentReply) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("comment_replies").
			Unique().
			Field("user_id").
			Required(),
		edge.From("comment", Comment.Type).
			Ref("comment_replies").
			Unique().
			Field("comment_id").
			Required(),
	}
}
