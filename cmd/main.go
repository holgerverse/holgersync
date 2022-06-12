package main

import (
	"log"

	"github.com/holgerverse/holgersync/commands"
	"github.com/holgerverse/holgersync/pkg/logger"
)

func main() {

	err := commands.Execute()
	if err != nil {
		log.Fatal(err)
	}

	// Create a new logger config
	loggerConfig := make(map[string]string)

	// Set the logger level
	if commands.Debug {
		loggerConfig["level"] = "debug"
	} else {
		loggerConfig["level"] = "error"
	}

	// Create a new logger and initialize it with given config
	logger := logger.NewCmdLogger()
	logger.InitLogger(map[string]string{
		"level": loggerConfig["level"],
	})

	logger.Info("Holgersync logging initialized")

}
