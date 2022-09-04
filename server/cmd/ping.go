package cmd

import (
	"fmt"

	"github.com/BeanWei/tingyu/g"
	"github.com/urfave/cli/v2"
)

var (
	Ping = &cli.Command{
		Name:  "ping",
		Usage: "check all service",
		Action: func(ctx *cli.Context) error {
			fmt.Println("Ping Meilisearch...")
			if info, err := g.Meili().Health(); err != nil {
				return err
			} else {
				g.Dump(info)
			}
			return nil
		},
	}
)
