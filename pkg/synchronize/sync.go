package synchronize

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/holgerverse/holgersync/config"
	"github.com/holgerverse/holgersync/pkg/helpers"
	"github.com/holgerverse/holgersync/pkg/logger"
)

// Entrypoint for the synchronize command
func Sync(cfg *config.Config) {

	logger := logger.NewCliLogger(cfg)
	logger.InitLogger()
	logger.Debug("Logger initialized")

	//Read the content of the root fileBuildTargetConfig
	sourceFileContent, err := helpers.GetAbsPathAndReadFile(cfg.HolgersyncConfig.SourceFileConfig.FilePath)
	if err != nil {
		logger.Fatal(err)
	}

	targets := &cfg.HolgersyncConfig.Targets
	for _, target := range *targets {

		logger.Debugf("Processing target: %s", target.Path)
		targetFilePath := fmt.Sprintf("%s/%s", target.Path, filepath.Base(cfg.HolgersyncConfig.SourceFileConfig.FilePath))

		if _, err := os.Stat(targetFilePath); errors.Is(err, os.ErrNotExist) {
			logger.Debugf("%s does not exist. Copying source content", targetFilePath)
			os.WriteFile(targetFilePath, sourceFileContent, 0644)
		}

		result, err := helpers.CompareWithSource(targetFilePath, sourceFileContent)
		if err != nil {
			logger.Fatal(err)
		}

		status, err := target.CheckFileStatusCode(targetFilePath)
		if err != nil {
			logger.Fatal(err)
		}

		switch *status {
		case git.Unmodified:
			logger.Debugf("%s is up to date.", targetFilePath)
		case git.Untracked, git.Modified:
			logger.Debugf("%s needs to be commited and pushed.", targetFilePath)
		}

		if !result {
			logger.Debugf("%s has changed. Updating", targetFilePath)
			os.WriteFile(targetFilePath, sourceFileContent, 0644)
		} else {
			logger.Debugf("%s is up to date", targetFilePath)
		}

		//Create a new branch for the target
		err = target.CreateHolgersyncBranch()
		if err != nil {
			logger.Fatal(err)
		}

		err = target.CommitAndPush(targetFilePath)
		if err != nil {
			logger.Fatal(err)
		}
	}
}
