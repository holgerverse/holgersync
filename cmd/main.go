package main

import (
	"log"
	"os"
	"path/filepath"

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
							//TODO: Move to helpers
							absolutePath, err := filepath.Abs(c.String("module-path"))
							if err != nil {
								log.Fatal(err)
							}
							holgerdocs(absolutePath, "terraform")
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
