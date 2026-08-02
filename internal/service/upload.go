package service

import (
	"os"

	"path/filepath"

	"github.com/berzz/pebbleFS/internal/chunk"
	"github.com/berzz/pebbleFS/internal/metadata"
	"github.com/google/uuid"
)
func (p *Pebble) Put(path string) error {

	info, err := os.Stat(path)
	if err != nil {
		return err
	}

	uploadID := uuid.New().String()

	if err := p.Metadata.CreateUpload(uploadID, info.Name()); err != nil {
		return err
	}

	fileID, err := p.Metadata.CreateFile(
		info.Name(),
		info.Size(),
	)

	if err != nil {
		return err
	}

	chunks, errs := chunk.Split(path)

	for c := range chunks {

		disk, err := p.Strategy.PickDisk(uint64(c.Size))
		if err != nil {
			return err
		}

		root := filepath.Join(
			disk.MountPath,
			".pebble",
			"chunks",
		)

		chunkPath, err := chunk.Write(root, c)
		if err != nil {
			return err
		}

		created, err := p.Metadata.UpsertChunk(metadata.Chunk{
			ID:             c.ID,
			DiskID:         disk.ID,
			Path:           chunkPath,
			Size:           c.Size,
			ReferenceCount: 1,
		})

		if err != nil {
			return err
		}

		if created {
			p.Storage.ReserveSpace(
				disk.ID,
				uint64(c.Size),
			)
		}

		err = p.Metadata.AddChunkToFile(
			fileID,
			c.ID,
			c.Index,
		)

		if err != nil {
			return err
		}
	}

	if err := <-errs; err != nil {
		return err
	}

	return p.Metadata.CompleteUpload(uploadID)
}