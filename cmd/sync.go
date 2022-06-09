package main

import (
	"context"
	"crypto/sha256"
	"fmt"
)

type contextKey string

const (
	contextSourceFileContext  contextKey = "sourceFileContent"
	contextSourceFileChecksum contextKey = "sourceFileChecksum"
)

func calcFileChecksum(configCtx context.Context) [32]byte {

	sum := sha256.Sum256(configCtx.Value(contextSourceFileContext).([]byte))

	return sum

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

	fmt.Print(configCtx.Value(contextSourceFileChecksum))

}
