package main

import (
	"log"

	"github.com/holgerverse/holgersync/commands"
)

func main() {

	err := commands.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
