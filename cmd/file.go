package cmd

import (
	"fmt"
	"strings"

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
		fsize, err := fileio.GetFileSize(strings.Join(args, " "))
		if err != nil {
			fmt.Printf("Error: %s.\n", err)
		} else {
			fmt.Printf("File Size: %v bits\n", fsize)
		}

		chosenChunk := fileio.ChunkFile(strings.Join(args, " "), fsize)
		fmt.Printf("Chunk Size: %v\n", chosenChunk)

	},
}
