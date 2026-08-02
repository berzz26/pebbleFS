package metadata

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type DB struct {
	conn *sql.DB
}

func New(path string) (*DB, error) {

	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	wrapped := &DB{
		conn: db,
	}

	if err := wrapped.initialize(); err != nil {
		return nil, err
	}

	return wrapped, nil

}

func (d *DB) Close() error {
	return d.conn.Close()
}

func (db *DB) GetChunk(id string) (*Chunk, error) {

	var chunk Chunk

	err := db.conn.QueryRow(
		`
		SELECT
			id,
			disk_id,
			path,
			size,
			reference_count
		FROM chunks
		WHERE id = ?
		`,
		id,
	).Scan(
		&chunk.ID,
		&chunk.DiskID,
		&chunk.Path,
		&chunk.Size,
		&chunk.ReferenceCount,
	)

	if err != nil {
		return nil, err
	}

	return &chunk, nil
}
