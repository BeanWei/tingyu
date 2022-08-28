package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
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
		field.JSON("content", []map[string]any{}).Comment("回复内容"),
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
	}
}
