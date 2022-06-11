package commands

import (
	"log"

	"github.com/holgerverse/holgersync/pkg/synchronize"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(syncCmd)
}

var (
	syncCmd = &cobra.Command{
		Use:   "sync",
		Short: "Synchronize files between everything via CLI",
		Long:  ``,
		Run:   commandSync,
	}
)

func commandSync(ccmd *cobra.Command, args []string) {

	if len(args) > 0 {
		synchronize.Sync(args[0])
	} else {
		log.Fatalf("%s", "Please provide a path to the config file.")
	}

}
