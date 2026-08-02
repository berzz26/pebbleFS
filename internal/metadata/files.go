package metadata

import "time"

func (db *DB) CreateFile(name string, size int64) (int64, error) {

	result, err := db.conn.Exec(
		`
		INSERT INTO files(filename,size,created_at)
		VALUES(?,?,?)
		`,
		name,
		size,
		time.Now(),
	)

	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (db *DB) GetFileByName(name string) (*File, error) {

	var file File

	err := db.conn.QueryRow(
		`
		SELECT id, filename, size, created_at
		FROM files
		WHERE filename = ?
		`,
		name,
	).Scan(
		&file.ID,
		&file.Name,
		&file.Size,
		&file.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &file, nil
}
