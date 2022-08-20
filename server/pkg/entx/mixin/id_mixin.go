package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/BeanWei/tingyu/pkg/idgen"
)

var _ ent.Mixin = (*ID)(nil)

type ID struct {
	mixin.Schema
}

func (ID) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			Unique().
			Immutable().
			DefaultFunc(func() int64 {
				return idgen.NewID()
			}),
	}
}
