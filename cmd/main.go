package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

// var args struct {
// 	ActionGroup string `arg:"positional" help:"Possible values: holgerdocs"`
// 	ModulePath  string `arg:"required" help:"Relative path to the Terraform module you want to perform actions on."`
// }

func main() {

	// arg.MustParse(&args)

	app := &cli.App{
		Name:  "holgersync",
		Usage: "Generate documentation for everything.",
		Commands: []*cli.Command{
			&cli.Command{
				Name:     "holgerdocs",
				Category: "cli-commands",
				Subcommands: []*cli.Command{
					&cli.Command{
						Name:      "terraform",
						HelpName:  "terraform",
						Category:  "holgerdocs",
						Usage:     "holgersync hoglerdocs terraform --module-path=<path>",
						UsageText: "Generate documentation for Terraform modules.",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "module-path",
								Usage:    "Relative path to the Terraform module you want to perform actions on.",
								Required: true,
							},
						},
						Action: func(c *cli.Context) error {
							holgerdocs(absolutePath(c.String("module-path")), "terraform")
							return nil
						},
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
