package fileio

import (
	"fmt"
	"strings"

	"github.com/Uhm-why/gusher/chunk"
)

func ChunkFile(args []string) {
	var err error
	fname := strings.Join(args, " ")

	fi := FileInfo{Name: fname}
	fi.Size, err = chunk.GetFileSize(fname)

	if err != nil {
		fmt.Printf("Error: %s.\n", err)
	} else {
		fmt.Printf("File Size: %v bits\n", fi.Size)
	}

	fi.ChunkSize = chunk.SelectChunkSize(fi.Size)
	fmt.Printf("Chunk Size: %v\n", fi.ChunkSize)

	fi.NumberOfChunks, err = chunk.SelectNumberOfChunks(fi.Size, fi.ChunkSize)

	if fi.NumberOfChunks > 1 {
		fi.HasChunks = true
	} else {
		fi.HasChunks = false
	}

	if err != nil {
		fmt.Printf("Error: %s. \n", err)
	} else {
		fmt.Printf("File will be divided into %v chunks\n", fi.NumberOfChunks)
	}

	fi.Hash, err = HashFile(fi.Name)

	if err != nil {
		fmt.Printf("Error: %s.", err)
	} else {
		fmt.Printf("File SHA256 Hash: %s\n", fi.Hash)
	}

}
