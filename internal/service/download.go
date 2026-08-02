package service

import (
	"os"

	"github.com/berzz/pebbleFS/internal/chunk"
)

func (p *Pebble) Get(filename, destination string) error {

	file, err := p.Metadata.GetFileByName(filename)
	if err != nil {
		return err
	}

	entries, err := p.Metadata.GetFileChunks(file.ID)
	if err != nil {
		return err
	}

	out, err := os.Create(destination)
	if err != nil {
		return err
	}

	defer out.Close()

	for _, entry := range entries {
		physical, err := p.Metadata.GetChunk(entry.ChunkID)
		if err != nil {
			return err
		}
		data, err := chunk.Read(physical.Path)
		if err != nil {
			return err
		}
		_, err = out.Write(data)
		if err != nil {
			return err
		}

	}

	return nil
}
