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
}

type Logger struct {
	Level       string
	Destination string
}

// Load the holgersync config file
func LoadConfig(filename string) *viper.Viper {

	v := viper.New()

	v.SetConfigFile("holgersync.yaml")
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		log.Fatal(err)
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
