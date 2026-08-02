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