package cmd

import (
	"github.com/BeanWei/tingyu/data/ent"
	"github.com/BeanWei/tingyu/data/ent/migrate"
	"github.com/urfave/cli/v2"
)

var (
	// TODO: Replace with Atlas Versioned-Migrations
	Migrate = &cli.Command{
		Name:  "migrate",
		Usage: "auto migrate database schema",
		Action: func(ctx *cli.Context) error {
			return ent.DB().Schema.Create(
				ctx.Context,
				migrate.WithDropIndex(true),
				migrate.WithDropColumn(true),
				migrate.WithForeignKeys(false),
			)
		},
	}
)
