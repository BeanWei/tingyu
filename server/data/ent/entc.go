//go:build ignore

package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/BeanWei/tingyu/pkg/entx"
)

func main() {
	err := entc.Generate("./schema",
		&gen.Config{
			Features: []gen.Feature{gen.FeatureVersionedMigration},
		},
		entc.Extensions(&entx.Extension{}),
	)
	if err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
