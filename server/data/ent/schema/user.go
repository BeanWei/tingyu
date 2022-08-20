package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/BeanWei/tingyu/pkg/entx/mixin"
)

type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.ID{},
		mixin.Time{},
	}
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Unique().Comment("用户名"),
		field.String("nickname").Default("").Comment("昵称"),
		field.String("password").Sensitive().Comment("Hash密码"),
		field.String("salt").Sensitive().Comment("盐值"),
		field.Int8("status").Default(0).Comment("状态(0.正常 1.停用)"),
		field.String("avatar").Default("").Comment("头像"),
		field.Bool("is_admin").Default(false).Comment("是否管理员"),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type),
		edge.To("comments", Comment.Type),
	}
}
