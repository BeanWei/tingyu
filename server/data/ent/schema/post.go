package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/BeanWei/tingyu/pkg/entx/mixin"
)

type Post struct {
	ent.Schema
}

func (Post) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.ID{},
		mixin.Time{},
	}
}

func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").Comment("用户ID"),
		field.Int("comment_count").Default(0).Comment("评论数"),
		field.Int8("visibility").Default(0).Comment("可见性(0.公开 1.私密 2.好友可见)"),
		field.Bool("is_top").Default(false).Comment("是否置顶"),
		field.Bool("is_excellent").Default(false).Comment("是否精华"),
		field.Bool("is_lock").Default(false).Comment("是否锁定"),
		field.Int64("latest_replied_at").Default(0).Comment("最后回复时间"),
		field.String("ip").Default("").Sensitive().Comment("IP"),
		field.String("ip_loc").Default("").Comment("IP地址"),
		field.String("content").Default("{}").Comment("帖子内容"),
	}
}

func (Post) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
	}
}

func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("posts").
			Unique().
			Field("user_id").
			Required(),
		edge.To("topics", Topic.Type),
		edge.To("comments", Comment.Type),
	}
}
