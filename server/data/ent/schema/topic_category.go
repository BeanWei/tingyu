package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/BeanWei/tingyu/pkg/entx/mixin"
)

type TopicCategory struct {
	ent.Schema
}

func (TopicCategory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.ID{},
		mixin.Time{},
	}
}

func (TopicCategory) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Default("").Comment("分类名称"),
		field.Int("rank").Default(9999).Comment("排序值"),
	}
}

func (TopicCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("topics", Topic.Type),
	}
}
