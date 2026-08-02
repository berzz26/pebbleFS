package placement

import (
	"errors"

	"github.com/berzz/pebbleFS/internal/storage"
)

// this module, we'll be depicting how which drive would get the first chunk from a pool of drives.
// to keep things simple for the v1, we'd be using first fit with most free space.

var ErrNoDiskAvailable = errors.New("no suitable disk available")

type Strategy interface {
	PickDisk(chunkSize uint64) (*storage.Disk, error)
}
