package chunk

import (
	"errors"
	"os"
	"path/filepath"
)

// ErrCorruptedChunk is returned when a chunk's data does not match its ID hash.
var ErrCorruptedChunk = errors.New("corrupted chunk")

func Write(root string, chunk Chunk) error {
	//rebuild the path from the chunk hash
	path := ChunkPath(root, chunk.ID)
	dir := filepath.Dir(path)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	//natural dedupe
	if _, err := os.Stat(path); err == nil {
		// Chunk already exists.
		return nil
	}
	
	if len(chunk.Data) == 0 {
		return errors.New("empty chunk")
	}
	computed := Hash(chunk.Data)

	if computed != chunk.ID {
		return ErrCorruptedChunk
	}

	return os.WriteFile(path, chunk.Data, 0644)
}
