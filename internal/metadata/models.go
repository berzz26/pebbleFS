package metadata

import "time"


type File struct {
	ID        int64
	Name      string
	Size      int64
	CreatedAt time.Time
}

type Chunk struct {
	ID             string
	DiskID         string
	Path           string
	Size           int64
	ReferenceCount int64
}

type FileChunk struct {
	FileID     int64
	ChunkID    string
	ChunkIndex int
}

type UploadStatus string

const (
	UploadWriting UploadStatus = "WRITING"
	UploadComplete UploadStatus = "COMPLETE"
	UploadFailed UploadStatus = "FAILED"
)

type Upload struct {
	ID        string
	FileName  string
	Status    UploadStatus
	CreatedAt time.Time
}

func (db *DB) initialize() error {

	query := `
CREATE TABLE IF NOT EXISTS uploads(
    id TEXT PRIMARY KEY,
    filename TEXT NOT NULL,
    status TEXT NOT NULL,
    created_at DATETIME NOT NULL
);

CREATE TABLE IF NOT EXISTS files(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    filename TEXT NOT NULL,
    size INTEGER NOT NULL,
    created_at DATETIME NOT NULL
);

CREATE TABLE IF NOT EXISTS chunks(
    id TEXT PRIMARY KEY,
    disk_id TEXT NOT NULL,
    path TEXT NOT NULL,
    size INTEGER NOT NULL,
    reference_count INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS file_chunks(
    file_id INTEGER NOT NULL,
    chunk_id TEXT NOT NULL,
    chunk_index INTEGER NOT NULL,

    PRIMARY KEY(file_id, chunk_index),

    FOREIGN KEY(file_id) REFERENCES files(id),
    FOREIGN KEY(chunk_id) REFERENCES chunks(id)
);
`

	_, err := db.conn.Exec(query)
	return err
}