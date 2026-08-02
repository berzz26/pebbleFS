package service

import (
	"fmt"

	"github.com/berzz/pebbleFS/internal/metadata"
	"github.com/berzz/pebbleFS/internal/placement"
	"github.com/berzz/pebbleFS/internal/storage"
)

type Pebble struct {
	Metadata *metadata.DB
	Storage  *storage.Manager
	Strategy placement.Strategy
}

func New() (*Pebble, error) {
	db, err := metadata.New("data/pebble.db")

	if err != nil {
		return nil, err
	}

	storageManager := storage.NewManager()
	disks, err := db.GetDisks()
	if err != nil {
		return nil, err
	}

	for _, disk := range disks {
		if err := storageManager.AddDisk(disk.MountPath); err != nil {
			fmt.Printf("warning: failed to load disk %s: %v\n", disk.ID, err)
		}
	}
	strategy := placement.NewMostFreeStrategy(storageManager)

	return &Pebble{
		Metadata: db,
		Storage:  storageManager,
		Strategy: strategy,
	}, nil
}
