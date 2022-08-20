package entx

import (
	"embed"
	"text/template"

	"entgo.io/ent/entc/gen"
)

var (
	//go:embed template
	templateDir embed.FS
	FuncMap     = template.FuncMap{}
	Templates   = gen.MustParse(gen.NewTemplate("lifly").Funcs(FuncMap).ParseFS(templateDir, "template/*tmpl"))
)
