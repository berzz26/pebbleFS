package service

import "github.com/berzz/pebbleFS/internal/chunk"

func (p *Pebble) Delete(name string) error {
	//find the file
	file, err := p.Metadata.GetFileByName(name)
	if err != nil {
		return err
	}

	entries, err := p.Metadata.GetFileChunks(file.ID)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		physical, err := p.Metadata.GetChunk(entry.ChunkID)
		if err != nil {
			return err
		}
		refs, err := p.Metadata.DecrementReference(physical.ID)
		if err != nil {
			return err
		}

		if refs == 0 {
			//refs zero means no copy exist, in that case, delete the physical chunk and release disk space.
			err = chunk.Delete(physical.Path)
			if err != nil {
				return err
			}

			err = p.Metadata.DeleteChunk(physical.ID)
			if err != nil {
				return err
			}

			p.Storage.ReleaseSpace(
				physical.DiskID,
				uint64(physical.Size),
			)
		}

	}
	// if refs are not zero, delete mappings and then delete the record
	err = p.Metadata.DeleteFileChunks(file.ID)
	if err != nil {
		return err
	}

	err = p.Metadata.DeleteFile(file.ID)
	if err != nil {
		return err
	}

	return nil

}
