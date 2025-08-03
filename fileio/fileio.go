package fileio

import (
	"fmt"
	"strings"

	"github.com/Uhm-why/gusher/chunk"
)

func ChunkFile(args []string) {
	fsize, err := chunk.GetFileSize(strings.Join(args, " "))

	if err != nil {
		fmt.Printf("Error: %s.\n", err)
	} else {
		fmt.Printf("File Size: %v bits\n", fsize)
	}

	var chunkSize int64 = chunk.SelectChunkSize(fsize)
	fmt.Printf("Chunk Size: %v\n", chunkSize)

	numberOfChunks, err := chunk.SelectNumberOfChunks(fsize, chunkSize)
	if err != nil {
		fmt.Printf("Error: %s. \n", err)
	} else {
		fmt.Printf("File will be divided into %v chunks\n", numberOfChunks)
	}

	fileHash, err := HashFile(strings.Join(args, ""))

	if err != nil {
		fmt.Printf("Error: %s.", err)
	} else {
		fmt.Printf("File SHA256 Hash: %s\n", fileHash)
	}

}
