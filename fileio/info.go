package fileio

import (
	"fmt"

	"github.com/Uhm-why/gusher/chunk"
)

type FileInfo struct {
	Name           string
	Size           int64
	Hash           string
	HasChunks      bool
	ChunkSize      int64
	NumberOfChunks int64
	Children       []string
}

type ChunkInfo struct {
	Name   string
	Size   int64
	Hash   string
	Parent string
}

func GetFileInfo(fname string) FileInfo {
	var err error

	var fi FileInfo = FileInfo{Name: fname}
	fi.Size, err = chunk.GetFileSize(fname)

	if err != nil {
		fmt.Printf("Error: %s.\n", err)
	} else {
		fmt.Printf("File Size: %v bits\n", fi.Size)
	}

	fi.ChunkSize = chunk.SelectChunkSize(fi.Size)
	fmt.Printf("Chunk Size: %v\n", fi.ChunkSize)

	fi.NumberOfChunks, err = chunk.SelectNumberOfChunks(fi.Size, fi.ChunkSize)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		if fi.NumberOfChunks > 1 {
			fi.HasChunks = true
			fmt.Printf("File will be divided into %v chunks\n", fi.NumberOfChunks)
		} else {
			fi.HasChunks = false
			fmt.Printf("File will be sent as one chunk.\n")
		}

	}

	fi.Hash, err = HashFile(fi.Name)

	if err != nil {
		fmt.Printf("Error: %s.\n", err)
	} else {
		fmt.Printf("File SHA256 Hash: %s\n", fi.Hash)
	}

	return fi
}
