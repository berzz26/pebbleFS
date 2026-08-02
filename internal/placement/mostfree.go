package placement

import (
	"github.com/berzz/pebbleFS/internal/storage"
)

type MostFreeStrategy struct {
	manager *storage.Manager
}

func NewMostFreeStrategy(m *storage.Manager) *MostFreeStrategy {
	return &MostFreeStrategy{
		manager: m,
	}
}

func (s *MostFreeStrategy) PickDisk(chunkSize uint64) (*storage.Disk, error) {

	var best *storage.Disk

	for _, disk := range s.manager.Disks() {

		//check health
		if !disk.Healthy {
			continue
		}
		//check free space
		if disk.FreeSpace < chunkSize {
			continue
		}
		//select the best choice
		if best == nil || disk.FreeSpace > best.FreeSpace {
			best = disk
		}
	}

	if best == nil {
		return nil, ErrNoDiskAvailable
	}

	return best, nil
}