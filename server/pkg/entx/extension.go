package entx

import (
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

var _ entc.Extension = (*Extension)(nil)

type Extension struct {
	entc.DefaultExtension
}

func (Extension) Templates() []*gen.Template {
	return []*gen.Template{Templates}
}
