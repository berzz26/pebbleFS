package service

import (
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
	strategy := placement.NewMostFreeStrategy(storageManager)

	return &Pebble{
		Metadata: db,
		Storage:  storageManager,
		Strategy: strategy,
	},nil
}
