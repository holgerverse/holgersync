package commands

import (
	"log"

	"github.com/holgerverse/holgersync/config"
	"github.com/holgerverse/holgersync/pkg/synchronize"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(syncCmd)
	syncCmd.Flags().StringVarP(&holgersyncConfigPath, "config-file", "c", "./holgersyncfile.yml", "Path to the holgersync config file")
}

var (
	holgersyncConfigPath string
	syncCmd              = &cobra.Command{
		Use:   "sync",
		Short: "Synchronize files between everything via CLI",
		Long:  ``,
		Run:   commandSync,
	}
)

func commandSync(ccmd *cobra.Command, args []string) {

	// Read the Holgersync configuration file
	cfgFile, err := config.LoadConfig(holgersyncConfigPath)
	if err != nil {
		log.Fatalln("No Holgersync configuration file found in this directory.")
	}

	// Parse Holgersync config
	cfg := config.ParseConfig(cfgFile)

	if Debug {
		cfg.Logger.Level = "debug"
	} else {
		cfg.Logger.Level = "error"
	}

	cfg.Logger.Destination = LogToFile

	synchronize.Sync(cfg)
}
