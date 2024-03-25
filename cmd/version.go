package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of iptracker",
	Long:  `Print the version number of iptracker`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version 1.20.00")
	},
}
