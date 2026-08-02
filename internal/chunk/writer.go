package chunk

import (
	"errors"
	"os"
	"path/filepath"
)

// ErrCorruptedChunk is returned when a chunk's data does not match its ID hash.
var ErrCorruptedChunk = errors.New("corrupted chunk")
var ErrEmptyChunk = errors.New("empty chunk")

func Write(root string, chunk Chunk) (string, error) {
	//rebuild the path from the chunk hash
	path := ChunkPath(root, chunk.ID)
	dir := filepath.Dir(path)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return "",err
	}
	//natural dedupe
	if _, err := os.Stat(path); err == nil {
		// Chunk already exists.
		return path,nil
	}

	if len(chunk.Data) == 0 {
		return "",ErrEmptyChunk
	}
	computed := Hash(chunk.Data)

	if computed != chunk.ID {
		return "",ErrCorruptedChunk
	}
	if err := os.WriteFile(path, chunk.Data, 0644); err != nil {
		return "", err
	}

	return path, nil
}
