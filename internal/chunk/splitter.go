package chunk

import (
	"io"
	"os"

	"github.com/google/uuid"

	"github.com/berzz/pebbleFS/internal/config"
)
// here, well be using stream processing with channels.
// instead of waiting for the ENTIRE file to be read before writing a single chunk, pebbleFs will be writing the chunk as soon as its read.
func Split(path string) (<-chan Chunk, <-chan error) {
	chunkChan := make(chan Chunk)
	errChan := make(chan error, 1)

	go func() {
		defer close(chunkChan)
		defer close(errChan)

		file, err := os.Open(path)
		if err != nil {
			errChan <- err
			return
		}
		defer file.Close()

		buffer := make([]byte, config.ChunkSize)
		index := 0

		for {
			n, err := file.Read(buffer)
			if err != nil {
				if err == io.EOF {
					break
				}
				errChan <- err
				return
			}
			/*
				file.Read() reuses the same buffer every iteration.

				If we directly send `buffer`,
				the next Read() would overwrite it.

				So we create a new slice and copy
				only the valid bytes.
			*/

			data := make([]byte, n)
			copy(data, buffer[:n])

			chunkChan <- Chunk{
				ID:    uuid.New().String(),
				Index: index,
				Data:  data,
				Size:  int64(n),
			}

			index++

		}
	}()
	return chunkChan, errChan
}
