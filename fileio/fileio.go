package fileio

import (
	"fmt"
	"os"
)

func FileExists(fname string) bool {
	_, err := os.Stat(fname)

	return !os.IsNotExist(err)
}

func ChunkFile(fi FileInfo) {

	if !FileExists(fi.Name) {
		fmt.Printf("File %s does not exist\n", fi.Name)
		return
	}

}
