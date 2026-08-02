package service

import (
	"github.com/berzz/pebblefs/internal/metadata"
)

type Pebble struct {
	Metadata *metadata.db
}

func New() (*Pebble, error) {
	db, err := metadata.new("data/pebble.db")
	if err != nil {
		return nil, err
	}

	return &Pebble{
		Metadata: db,
	}, nil
}
