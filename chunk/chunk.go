package chunk

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

func SelectChunkSize(fsize int64) int64 {
	if fsize <= 0 {
		return 0
	}
	if fsize >= maxChunkSize {
		return maxChunkSize
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

func SelectNumberOfChunks(fsize int64, csize int64) (int64, error) {
	switch {
	case fsize <= 0 || csize <= 0:
		return 0, errors.New("error: no file size and/or chunk size recieved at SelectNumberOfChunks")
	case fsize%csize == 0:
		return fsize / csize, nil
	case fsize%csize > 0:
		return (fsize / csize) + 1, nil
	default:
		return fsize / csize, nil
	}

}
