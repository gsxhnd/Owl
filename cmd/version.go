package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var versionCmd = &cli.Command{
	Name:  "version",
	Usage: "show version info",
	Action: func(ctx *cli.Context) error {
		fmt.Println("version")
		return nil
	},
}
