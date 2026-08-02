package storage
import "time"

type DiskMetadata struct {
	ID         string    `json:"id"`
	Version    int       `json:"version"`
	CreatedAt  time.Time `json:"created_at"`
	Filesystem string    `json:"filesystem"`
	Label      string    `json:"label"`
}