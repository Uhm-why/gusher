package cmd

import (
	"fmt"

	"github.com/Uhm-why/gusher/fileio"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(fileCmd)
}

var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "File to Send.",
	Long:  "State the name of the file that you would like to send.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Reading File: %s\n", args)
		fileio.ChunkFile(args)

	},
}
