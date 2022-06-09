package main

import (
	"log"

	"gopkg.in/yaml.v2"
)

type holgersyncConfig struct {
	Config struct {
		SourceFile string `yaml:"sourceFile"`
		RootPath   string `yaml:"rootPath"`
		Depth      int64
		File       string
	}
}

func (c *holgersyncConfig) readConfig(path string) *holgersyncConfig {

	configFile := getAbsPathAndReadFile(path)

	err := yaml.Unmarshal(configFile, c)
	if err != nil {
		log.Fatal(err)
	}

	return c
}
