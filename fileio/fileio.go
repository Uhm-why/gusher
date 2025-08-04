package fileio

import (
	"fmt"
	"strings"
)

func ChunkFile(args []string) {
	fi := GetFileInfo(strings.Join(args, " "))
	fmt.Printf("TODO: Implement File Chunking Logic... %v\n", fi.Name)

}
