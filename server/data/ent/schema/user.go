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
		field.Int8("status").Default(0).Comment("状态(1.正常 2.停用)"),
		field.String("avatar").Default("").Comment("头像"),
		field.String("headline").Default("").Comment("个人简介"),
		field.Bool("is_admin").Default(false).Comment("是否管理员"),
		field.Int("count_post").Default(0).Comment("发贴数量").StructTag(`json:"count_post"`),
		field.Int("count_topic").Default(0).Comment("关注话题数量").StructTag(`json:"count_topic"`),
		field.Int("count_following").Default(0).Comment("关注数量").StructTag(`json:"count_following"`),
		field.Int("count_follower").Default(0).Comment("粉丝数量").StructTag(`json:"count_follower"`),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("topics", Topic.Type),
		edge.To("posts", Post.Type),
		edge.To("comments", Comment.Type),
		edge.To("comment_replies", CommentReply.Type),
		edge.To("user_reactions", UserReaction.Type),
	}
}
