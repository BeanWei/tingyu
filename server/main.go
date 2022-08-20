package main

import (
	"os"

	"github.com/BeanWei/tingyu/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "LiFly",
		Commands: []*cli.Command{
			cmd.Start,
		},
	}
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
