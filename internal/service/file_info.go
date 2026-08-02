package service
import "github.com/berzz/pebbleFS/internal/metadata"

func (p *Pebble) Info(name string) (*metadata.File, []metadata.ChunkInfo, error) {

	file, err := p.Metadata.GetFileByName(name)
	if err != nil {
		return nil, nil, err
	}

	chunks, err := p.Metadata.GetChunkInfo(file.ID)
	if err != nil {
		return nil, nil, err
	}

	return file, chunks, nil
}