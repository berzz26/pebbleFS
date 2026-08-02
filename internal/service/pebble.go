package service

import (
	"github.com/berzz/pebbleFS/internal/metadata"
)

type Pebble struct {
	Metadata *metadata.DB
}

func New() (*Pebble, error) {
	db, err := metadata.New("data/pebble.db")
	if err != nil {
		return nil, err
	}

	return &Pebble{
		Metadata: db,
	}, nil
}
