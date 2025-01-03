package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "owl"
	app.Commands = []*cli.Command{
		serverCmd,
		versionCmd,
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
