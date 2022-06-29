package synchronize

import (
	"errors"
	"fmt"
	"log"
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

	// Read the content of the root file
	sourceFileContent, err := helpers.GetAbsPathAndReadFile(cfg.HolgersyncConfig.SourceFileConfig.FilePath)
	if err != nil {
		logger.Fatal(err)
	}

	// Iterate over all targets
	for _, target := range cfg.HolgersyncConfig.Targets {

		logger.Debugf("Processing target: %s", target.Path)
		targetFilePath := fmt.Sprintf("%s/%s", target.Path, filepath.Base(cfg.HolgersyncConfig.SourceFileConfig.FilePath))

		if _, err := os.Stat(target.Path + "/" + filepath.Base(cfg.HolgersyncConfig.SourceFileConfig.FilePath)); errors.Is(err, os.ErrNotExist) {
			logger.Debugf("%s does not exist. Copying source content", targetFilePath)
			os.WriteFile(targetFilePath, sourceFileContent, 0644)
		}

		targetContent, err := helpers.GetAbsPathAndReadFile(targetFilePath)
		if err != nil {
			logger.Fatal(err)
		}

		result, err := helpers.CompareData(sourceFileContent, targetContent)
		if err != nil {
			log.Fatal(err)
		}
		if !result {
			logger.Debugf("%s has changed. Updating", targetFilePath)
			os.WriteFile(targetFilePath, sourceFileContent, 0644)
		}
	}
}
