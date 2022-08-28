package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/BeanWei/tingyu/pkg/entx/mixin"
)

type Topic struct {
	ent.Schema
}

func (Topic) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.ID{},
		mixin.Time{},
	}
}

func (Topic) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty().Default("").Comment("标题"),
		field.String("icon").Default("").Comment("图标"),
		field.String("description").Default("").Comment("描述"),
		field.String("notice").Default("").Comment("声明"),
		field.Int("post_count").Default(0).Comment("帖子数量"),
		field.Int("follower_count").Default(0).Comment("关注数量"),
		field.Int("attender_count").Default(0).Comment("参与者数量"),
	}
}

func (Topic) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("posts", Post.Type).Ref("topics"),
	}
}
