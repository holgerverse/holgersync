package main

import (
	"log"

	"github.com/holgerverse/holgersync/commands"
	"github.com/holgerverse/holgersync/pkg/logger"
)

func main() {

	logger.NewZapLogger()

	err := commands.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
