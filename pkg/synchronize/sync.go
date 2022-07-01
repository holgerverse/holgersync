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
	"github.com/holgerverse/holgersync/pkg/remotes"
)

// Entrypoint for the synchronize command
func Sync(cfg *config.Config) {

	logger := logger.NewCliLogger(cfg)
	logger.InitLogger()
	logger.Debug("Logger initialized")

	// Read the content of the root fileBuildTargetConfig
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
			logger.Fatal(err)
		}

		remote, err := remotes.CheckFileStatusCode(target.Path, filepath.Base(cfg.HolgersyncConfig.SourceFileConfig.FilePath))
		if err != nil {
			logger.Fatal(err)
		}

		if result {
			logger.Debugf("%s is up to date", targetFilePath)
			break
		}

		if *remote == git.Unmodified {
			logger.Debugf("%s is umodified.", targetFilePath)
			break
		}

		logger.Debugf("%s has changed. Updating", targetFilePath)
		os.WriteFile(targetFilePath, sourceFileContent, 0644)

		err = remotes.CreateNewBranch(target.Path)
		if err != nil {
			logger.Fatal(err)
		}

		err = remotes.CommitAndPush(target.Path, filepath.Base(targetFilePath))
		if err != nil {
			logger.Fatal(err)
		}

	}
}
