package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {

	// Create Loglevels
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})

	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
		zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority),
	)

	logger := zap.New(core, zap.WithCaller(true), zap.OnFatal(zapcore.CheckWriteAction(3)))
	defer logger.Sync()

	logger.Info("Logger initiliazed.")

	app := &cli.App{
		Name:  "holgersync",
		Usage: "Generate documentation for everything.",
		Commands: []*cli.Command{{
			Name:     "holgerdocs",
			Category: "cli-commands",
			Subcommands: []*cli.Command{{
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
