package fileio

import (
	"fmt"
	"io"
	"os"
)

func FileExists(fname string) bool {
	_, err := os.Stat(fname)

	return !os.IsNotExist(err)
}

func ProcessFileWithChunks(fi FileInfo) error {
	sourceF, err := os.Open(fi.Name)
	ci := ChunkInfo{Parent: fi.Name}

	if err != nil {
		return fmt.Errorf("error opening source file: %v", err)
	}

	defer sourceF.Close()

	fi.Children = make([]ChunkInfo, 0, fi.NumberOfChunks)

	for i := int64(0); i < fi.NumberOfChunks; i++ {

		ci.Name = fmt.Sprintf("%s.chunk%d", fi.Name, i)
		chunkFile, err := os.Create(ci.Name)

		if err != nil {
			return fmt.Errorf("error creating chunk %d: %v", i, err)
		}

		ci.Size = fi.ChunkSize

		if i == fi.NumberOfChunks-1 {
			ci.Size = fi.Size - (i * fi.ChunkSize)
		}

		thisChunkFile, err := io.CopyN(chunkFile, sourceF, ci.Size)
		chunkFile.Close()

		if err != nil {
			return fmt.Errorf("error writing to chunk %d: %v", i, err)
		}

		fmt.Printf("Wrote chunk %d: (%d bytes)\n", i+1, thisChunkFile)

		ci.Hash, err = HashFile(ci.Name)

		fi.Children = append(fi.Children, ci)

		if err != nil {
			return fmt.Errorf("error hashing chunk %d: %v", i, err)
		}

		fmt.Printf("Chunk Hash: %v\n", ci.Hash)

	}

	return nil
}

func ChunkFile(fi FileInfo) {

	if !FileExists(fi.Name) {
		fmt.Printf("File %s does not exist\n", fi.Name)
		return
	}

	if fi.HasChunks {
		ProcessFileWithChunks(fi)
	}

}
