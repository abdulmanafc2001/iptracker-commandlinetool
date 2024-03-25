package cmd

import "github.com/spf13/cobra"

var (
	rootCmd = &cobra.Command{
		Use:     "iptracker",
		Short:   "Cli tool for iptracking",
		Long:    "CLI tool for iptracking",
		Example: "iptracker check http://example.com",
	}
)

func Execute() error {
	return rootCmd.Execute()
}
