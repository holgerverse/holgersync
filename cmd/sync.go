package main

import (
	"context"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type contextKey string

const (
	contextSourceFileContent  contextKey = "sourceFileContent"
	contextSourceFileChecksum contextKey = "sourceFileChecksum"
)

func updateFile(rootPath string, path string, ctx context.Context) error {

	in, err := os.Open(path)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	stringReader := strings.NewReader(ctx.Value(contextSourceFileContent).(string))

	_, err = io.Copy(out, stringReader)
	if err != nil {
		return err
	}

	return out.Close()

}

func getPaths(rootPath string, fileRegex string, ctx context.Context) error {

	// Check if the root path is exists
	_, err := os.Stat(rootPath)
	if os.IsNotExist(err) {
		log.Fatal("Root path does not exist.")
	}

	// Walk through the root path
	err = filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {

		// Catch errors
		if err != nil {
			return err
		}

		r, err := regexp.Compile(fileRegex)
		if err != nil {
			log.Fatal(err)
		}

		if !r.MatchString(info.Name()) {
			return err
		}

		//Calculate the checksum of the file
		checksum := calcFileChecksum(path)

		if checksum != ctx.Value(contextSourceFileChecksum) {
			log.Printf("File %s does not match the source file.\n", path)
			err = updateFile(rootPath, path, ctx)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return err
}

func sync(path string) {

	// Create a new hoglersyncConfig object
	config := &holgersyncConfig{}

	// Create a context for holgersync configuration settings
	configCtx := context.TODO()

	// Read the content of the given holgersync config file
	config.readConfig(path)

	configCtx = context.WithValue(configCtx, contextSourceFileContent, getAbsPathAndReadFile(config.Config.SourceFile))
	configCtx = context.WithValue(configCtx, contextSourceFileChecksum, calcFileChecksum(config.Config.SourceFile))

	err := getPaths(config.Config.RootPath, config.Config.FileRegex, configCtx)
	if err != nil {
		log.Fatal(err)
	}

}
