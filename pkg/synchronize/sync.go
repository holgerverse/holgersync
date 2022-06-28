package synchronize

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

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

		// Create new branch
		err = remotes.CreateNewBranch(target.Path)
		if err != nil {
			logger.Fatal(err)
		}

		err = remotes.CommitAndPush(target.Path, "test.json")
		if err != nil {
			logger.Fatal(err)
		}

	}
}
