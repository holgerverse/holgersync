package synchronize

import (
	"github.com/holgerverse/holgersync/pkg/helpers"
	"gopkg.in/yaml.v2"
)

type holgersyncConfig struct {
	Config struct {
		SourceFile string `yaml:"sourceFile"`
		RootPath   string `yaml:"rootPath"`
		FileRegex  string `yaml:"fileRegex"`
	}
}

func (c *holgersyncConfig) readConfig(path string) (*holgersyncConfig, error) {

	configFile, err := helpers.GetAbsPathAndReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal([]byte(configFile), c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
