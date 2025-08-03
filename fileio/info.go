package fileio

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
