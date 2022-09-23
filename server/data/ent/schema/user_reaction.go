package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/BeanWei/tingyu/pkg/entx/mixin"
)

type UserReaction struct {
	ent.Schema
}

func (UserReaction) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.ID{},
	}
}

func (UserReaction) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("created_at").Immutable().DefaultFunc(func() int64 {
			return time.Now().Unix()
		}),
		field.Int64("user_id").Comment("用户ID"),
		field.Enum("subject_type").Values("post", "comment", "reply").Comment("对象类型"),
		field.Int64("subject_id").Comment("对象ID"),
		field.Enum("react_code").Values(
			"emoji-thumbs_up", "emoji-thumbs_down", "emoji-laugh",
			"emoji-hooray", "emoji-confused", "emoji-heart",
			"emoji-rocket", "emoji-eyes",
			"action-star", "action-pin",
		).Comment("反应码"),
	}
}

func (UserReaction) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at"),
		index.Fields("user_id"),
		index.Fields("subject_type"),
		index.Fields("subject_id"),
		index.Fields("react_code"),
	}
}

func (UserReaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("user_reactions").
			Unique().
			Field("user_id").
			Required(),
	}
}
