package synchronize

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/holgerverse/holgersync/config"
	"github.com/holgerverse/holgersync/pkg/helpers"
	"github.com/holgerverse/holgersync/pkg/logger"
	"github.com/holgerverse/holgersync/pkg/remotes"
)

type contextKey string

const (
	contextSourceFileContent  contextKey = "sourceFileContent"
	contextSourceFileChecksum contextKey = "sourceFileChecksum"
)

func pushToBackend(path string) error {

	err := remotes.CreateNewBranch(path)
	if err != nil {
		return err
	}

	err = remotes.CommitAndPush(path, "test.json")
	if err != nil {
		return err
	}

	return nil

}

// Entrypoint for the synchronize command
func Sync(cfg *config.Config) {

	logger := logger.NewCliLogger(cfg)
	logger.InitLogger()
	logger.Debug("Logger initialized")

	configCtx := context.TODO()

	// Read the content of the root file
	sourceFileContent, err := helpers.GetAbsPathAndReadFile(cfg.HolgersyncConfig.SourceFileConfig.FilePath)
	if err != nil {
		logger.Fatal(err)
	}
	configCtx = context.WithValue(configCtx, contextSourceFileContent, sourceFileContent)

	sourceFileChecksum, err := helpers.CalcFileChecksum(sourceFileContent)
	if err != nil {
		logger.Fatal(err)
	}
	configCtx = context.WithValue(configCtx, contextSourceFileChecksum, sourceFileChecksum)

	// Iterate over all targets
	for _, target := range cfg.HolgersyncConfig.Targets {

		logger.Debugf("Processing target: %s", target.Path)
		targetFilePath := fmt.Sprintf("%s/%s", target.Path, filepath.Base(cfg.HolgersyncConfig.SourceFileConfig.FilePath))

		if _, err := os.Stat(target.Path + "/" + filepath.Base(cfg.HolgersyncConfig.SourceFileConfig.FilePath)); errors.Is(err, os.ErrNotExist) {
			logger.Debugf("%s does not exist. Copying source content", targetFilePath)
			os.WriteFile(targetFilePath, sourceFileContent, 0644)
		}

		sourceSha256, err := helpers.CalcFileChecksum(sourceFileContent)
		if err != nil {
			logger.Fatal(err)
		}

		targetContent, err := helpers.GetAbsPathAndReadFile(targetFilePath)
		if err != nil {
			logger.Fatal(err)
		}

		targetSha256, err := helpers.CalcFileChecksum(targetContent)
		if err != nil {
			logger.Fatal(err)
		}

		res := bytes.Compare(sourceSha256, targetSha256)
		if res != 0 {
			logger.Debugf("%s has changed. Updating", targetFilePath)
			os.WriteFile(targetFilePath, sourceFileContent, 0644)
		}

		fmt.Println(filepath.Base(targetFilePath))

		fmt.Println(os.Stat(targetFilePath))

	}
}
