package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type config struct {
	HolgerSyncConfig struct {
	}
}

func configLoader() {

	// Get the path of the home director.
	homePath, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	// Create path to config directory.
	configPath := homePath + "/.holgersync"

	// Check if config folder is created.
	_, err = os.Stat(configPath)
	if os.IsNotExist(err) {
		os.Mkdir(configPath, 0777)
	}

	// Append config file name.
	configPath = configPath + "/config.yml"

	// Check if config file exists, if not create a new one.
	_, err = os.Stat(configPath)
	if os.IsNotExist(err) {
		os.Create(configPath)
	}

}

func readConfig(filePath string) {

	configFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(configFile)

}
