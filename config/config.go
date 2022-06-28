package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	HolgersyncConfig HolgersyncConfig
	Logger           Logger
}

type HolgersyncConfig struct {
	SourceFileConfig SourceFileConfig `yaml:"sourceFileConfig"`
	Targets          []Target         `yaml:"Targets,mapstructure"`
}

type SourceFileConfig struct {
	FilePath  string `yaml:"filePath"`
	RootPath  string `yaml:"rootPath"`
	FileRegex string `yaml:"fileRegex"`
}

type Target struct {
	Path       string      `yaml:"path"`
	Parameters []Parameter `yaml:"parameters,mapstructure"`
}

type Parameter struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

type Logger struct {
	Level       string
	Destination string
}

// Load the holgersync config file
func LoadConfig(filename string) (*viper.Viper, error) {

	v := viper.New()

	v.SetConfigFile(filename)
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	return v, nil
}

// TODO
func ParseConfig(v *viper.Viper) *Config {

	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	return &c
}
