package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/BeanWei/tingyu/pkg/entx/mixin"
)

type File struct {
	ent.Schema
}

func (File) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.ID{},
		mixin.Time{},
	}
}

func (File) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").Comment("上传者ID"),
		field.Int64("size").Default(0).Optional().Comment("文件大小"),
		field.String("type").Default("").Optional().Comment("文件类型"),
		field.String("name").Default("").Optional().Comment("文件名"),
		field.String("orig_name").Default("").Optional().Comment("原始文件名"),
		field.String("url").Default("").Optional().Comment("文件链接"),
		field.String("oss_type").Default("").Optional().Comment("OSS类别"),
		field.String("oss_domain").Default("").Optional().Comment("OSS域名"),
		field.String("oss_bucket").Default("").Optional().Comment("OSS存储桶"),
	}
}

func (File) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("name"),
	}
}
