package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type config struct {
	HolgerSyncConfig struct {
	}
}

/*
	Reads the Holgerfile and returns the config struct

	Input: The path to the directory containing the Holgerfile

	Output: The config struct
*/

func loadConfig(directoryPath string) {

	// Get the path of the current directory.
	pwd, _ := os.Getwd()

	// Get the path to the directory containing the Holgerfile. Concatenate the current directory with the directoryPath.
	holgerfile, err := ioutil.ReadFile(filepath.Join(pwd, directoryPath))
	if err != nil {
		log.Fatal("Could not read Holgerfile in Directory: " + filepath.Join(pwd, directoryPath))
	}

	fmt.Print(holgerfile)

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
