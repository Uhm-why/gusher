package cmd

import (
	"fmt"
	"strings"

	"github.com/Uhm-why/gusher/fileio"
	"github.com/Uhm-why/gusher/store"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(chunkCmd)
	rootCmd.AddCommand(reintCmd)
}

var chunkCmd = &cobra.Command{
	Use:   "chunk",
	Short: "File to Chunk.",
	Long:  "State the name of the file that you would like to chunk.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Reading File: %s\n", args)

		fi := fileio.GetFileInfo(strings.Join(args, " "))

		fileio.ChunkFile(fi)

		store.SaveFileInfo(fi)

	},
}

var reintCmd = &cobra.Command{
	Use:   "reint",
	Short: "File to Reintegrate from chunks.",
	Long:  "State the name of the chunks (i.e. filename.chunk0 to reintegrate into the orgininal file)",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("Gathering chunks of file: %s\n", args)

	},
}
