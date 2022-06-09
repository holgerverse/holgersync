package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name:  "holgersync",
		Usage: "Synchronize files between everything",
		Commands: []*cli.Command{{
			Name:  "sync",
			Usage: "Synchronize based on configuration",
			Action: func(c *cli.Context) error {
				sync("tests/holgersyncfile.yml")
				return nil
			},
		}},
	}

	app.Run(os.Args)

}
