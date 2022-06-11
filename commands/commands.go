package commands

import "github.com/spf13/cobra"

var (
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
