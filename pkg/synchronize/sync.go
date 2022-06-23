package synchronize

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/holgerverse/holgersync/config"
	"github.com/holgerverse/holgersync/pkg/helpers"
	"github.com/holgerverse/holgersync/pkg/logger"
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

func GetPaths(rootPath string, fileRegex string, ctx context.Context) error {

	// Check if the root path is exists
	_, err := os.Stat(rootPath)
	if os.IsNotExist(err) {
		return err
	}

	// Walk through the root path
	err = filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {

		// Catch errors
		if err != nil {
			return err
		}

		r, err := regexp.Compile(fileRegex)
		if err != nil {
			return err
		}

		if !r.MatchString(info.Name()) {
			return err
		}

		// if checksum != ctx.Value(contextSourceFileChecksum) {
		// 	log.Printf("File %s does not match the source file.\n", path)
		// 	err = updateFile(rootPath, path, ctx)
		// 	if err != nil {
		// 		return err
		// 	}
		// }

		return nil
	})

	return err
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
}
