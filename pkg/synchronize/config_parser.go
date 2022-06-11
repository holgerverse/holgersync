package synchronize

import (
	"log"

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

func (c *holgersyncConfig) readConfig(path string) *holgersyncConfig {

	configFile := helpers.GetAbsPathAndReadFile(path)

	err := yaml.Unmarshal([]byte(configFile), c)
	if err != nil {
		log.Fatal(err)
	}

	return c
}
