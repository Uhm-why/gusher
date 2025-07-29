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
	Short: "Print Version Info",
	Long:  "This command prints the current version of gusher",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Gusher Infra Management v0.1")
	},
}
