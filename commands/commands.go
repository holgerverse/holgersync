package commands

import "github.com/spf13/cobra"

func init() {
	rootCmd.Flags().BoolVar(&Debug, "debug", false, "Enable debug mode")
}

var (
	Debug   bool
	rootCmd = &cobra.Command{
		Use:           "holgersync",
		Short:         "holgersync - Synchronize files between everything",
		Long:          ``,
		SilenceErrors: true,
		SilenceUsage:  true,
	}
)

func Execute() error {
	return rootCmd.Execute()
}
