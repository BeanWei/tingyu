package main

import (
	"os"

	"github.com/BeanWei/tingyu/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "TingYu",
		Commands: []*cli.Command{
			cmd.Start,
			cmd.Migrate,
			cmd.Ping,
		},
	}
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
