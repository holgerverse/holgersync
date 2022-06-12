package commands

import "github.com/spf13/cobra"

func init() {
	rootCmd.PersistentFlags().BoolVar(&Debug, "debug", false, "Enable debug mode")
	rootCmd.PersistentFlags().StringVar(&LogToFile, "log-to-file", "", "Safe logs to file")
}

var (
	Debug     bool
	LogToFile string
	rootCmd   = &cobra.Command{
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
