package main

import (
	"github.com/gsxhnd/owl/server/di"
	"github.com/urfave/cli/v2"
)

var serverCmd = &cli.Command{
	Name: "server",
	Action: func(ctx *cli.Context) error {
		app, err := di.InitApp()
		if err != nil {
			return err
		}
		return app.Run()
	},
}
