package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
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
		field.Int64("creator_id").Default(0).Comment("创建者ID"),
		field.Int("post_count").Default(0).Comment("帖子数量"),
		field.Int("follower_count").Default(0).Comment("关注数量"),
		field.Int("attender_count").Default(0).Comment("参与者数量"),
		field.Int64("topic_category_id").Optional().Default(0).Comment("分类ID"),
		field.Bool("is_rec").Default(false).Comment("是否推荐"),
		field.Int("rec_rank").Default(9999).Comment("推荐值"),
		// field.Bool("is_offline").Default(false).Comment("是否下线"),
		// field.Bool("is_unpostable").Default(false).Comment("禁止发贴"),
		// field.Bool("is_uncommentable").Default(false).Comment("禁止评论"),
	}
}

func (Topic) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("topic_category_id"),
		index.Fields("is_rec"),
	}
}

func (Topic) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("posts", Post.Type).Ref("topics"),
		edge.From("users", User.Type).Ref("topics"),
		edge.From("topic_categories", TopicCategory.Type).
			Ref("topics").
			Unique().
			Field("topic_category_id"),
	}
}
