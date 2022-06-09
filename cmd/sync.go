package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

type contextKey string

const (
	contextSourceFileContext  contextKey = "sourceFileContent"
	contextSourceFileChecksum contextKey = "sourceFileChecksum"
)

func getPaths(rootPath string, fileRegex string) {

	_, err := os.Stat(rootPath)
	if os.IsNotExist(err) {

		log.Fatal("Root path does not exist.")
	}

	err = filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		r, err := regexp.Compile(fileRegex)
		if err != nil {
			log.Fatal(err)
		}

		if r.MatchString(info.Name()) {
			fmt.Print(info.Name())
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

}

func sync(path string) {

	// Create a new hoglersyncConfig object
	config := &holgersyncConfig{}

	// Create a context for holgersync configuration settings
	configCtx := context.TODO()

	// Read the content of the given holgersync config file
	config.readConfig(path)

	configCtx = context.WithValue(configCtx, contextSourceFileContext, getAbsPathAndReadFile(config.Config.SourceFile))
	configCtx = context.WithValue(configCtx, contextSourceFileChecksum, calcFileChecksum(configCtx))

	getPaths(config.Config.RootPath, config.Config.FileRegex)

}
