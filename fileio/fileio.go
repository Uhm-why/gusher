package fileio

import (
	"errors"
	"os"
)

const (
	minChunkSize int64 = 4 << 10
	maxChunkSize int64 = 256 << 10
)

var chunkSizes = []int64{
	4 << 10,
	8 << 10,
	16 << 10,
	32 << 10,
	64 << 10,
	128 << 10,
	256 << 10,
}

func GetFileSize(fname string) (int64, error) {
	fileI, err := os.Stat(fname)
	if err != nil {
		return 0, errors.New("could not get file size. pre-gfs-return")
	}

	fileS := fileI.Size()

	return fileS, nil
}

func ChunkFile(fname string, fsize int64) int64 {

	chunkSize := SelectChunkSize(fsize)

	return chunkSize
}

func SelectChunkSize(fsize int64) int64 {
	if fsize <= 0 {
		return 0
	}

	if fsize <= minChunkSize {
		return fsize
	}

	for i := len(chunkSizes) - 1; i >= 0; i-- {
		if chunkSizes[i] <= fsize {
			return chunkSizes[i]
		}
	}

	return minChunkSize
}
