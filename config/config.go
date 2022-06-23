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
	TestingConfig TestingConfig `yaml:"testingConfig"`
	Cool          string
}

type TestingConfig struct {
	SourceFile string `yaml:"sourceFile"`
	RootPath   string `yaml:"rootPath"`
	FileRegex  string `yaml:"fileRegex"`
}

type Logger struct {
	Level       string
	Destination string
}

// Load the holgersync config file
func LoadConfig(filename string) *viper.Viper {

	v := viper.New()

	v.SetConfigFile(filename)
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %v", err)
	}

	return v
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
