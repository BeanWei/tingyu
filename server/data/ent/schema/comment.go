package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/BeanWei/tingyu/pkg/entx/mixin"
)

type Comment struct {
	ent.Schema
}

func (Comment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.ID{},
		mixin.Time{},
	}
}

func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("post_id").Comment("帖子ID"),
		field.Int64("user_id").Comment("用户ID"),
		field.String("ip").Default("").Sensitive().Comment("IP"),
		field.String("ip_loc").Default("").Comment("IP地址"),
		field.String("content").Default("").Comment("评论内容"),
		field.Int("reply_count").Default(0).Comment("回复数").StructTag(`json:"reply_count"`),
		field.Bool("is_poster").Default(false).Comment("是否发帖者"),
		field.Int8("status").Default(0).Comment("状态(1.审核通过 2.审核未通过 3.待审核)"),
	}
}

func (Comment) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("post_id"),
		index.Fields("user_id"),
	}
}

func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("comments").
			Unique().
			Field("user_id").
			Required(),
		edge.From("post", Post.Type).
			Ref("comments").
			Unique().
			Field("post_id").
			Required(),
		edge.To("comment_replies", CommentReply.Type),
	}
}
