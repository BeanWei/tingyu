package mixin

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

var _ ent.Mixin = (*Time)(nil)

type Time struct {
	mixin.Schema
}

func (Time) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("created_at").
			Immutable().
			DefaultFunc(func() int64 {
				return time.Now().Unix()
			}),
		field.Int64("updated_at").
			DefaultFunc(func() int64 {
				return time.Now().Unix()
			}).
			UpdateDefault(func() int64 {
				return time.Now().Unix()
			}),
		field.Int64("deleted_at").
			Optional().
			Default(0).
			StructTag(`json:"-"`),
	}
}

func (Time) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at"),
		index.Fields("deleted_at"),
	}
}
