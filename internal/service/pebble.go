package service

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/berzz/pebbleFS/internal/metadata"
	"github.com/berzz/pebbleFS/internal/placement"
	"github.com/berzz/pebbleFS/internal/storage"
)

type Pebble struct {
	Metadata *metadata.DB
	Storage  *storage.Manager
	Strategy placement.Strategy
}

func (p *Pebble) ListFiles() ([]metadata.File, error) {
	return p.Metadata.ListFiles()
}

func New() (*Pebble, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	dataDir := filepath.Join(home, ".pebble")
	dbPath := filepath.Join(dataDir, "pebble.db")

	db, err := metadata.New(dbPath)
	if err != nil {
		return nil, err
	}

	storageManager := storage.NewManager()
	disks, err := db.GetDisks()
	if err != nil {
		return nil, err
	}

	for _, disk := range disks {
		if _, err := storageManager.AddDisk(disk.MountPath); err != nil {
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
