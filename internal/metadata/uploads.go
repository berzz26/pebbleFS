package metadata

import "time"

func (db *DB) CreateUpload(id, filename string) error {
	_, err := db.conn.Exec(
		`
		INSERT INTO uploads(id, filename, status, created_at)
		VALUES(?,?,?,?)
		`,
		id,
		filename,
		UploadWriting,
		time.Now(),
	)

	return err

}

func (db *DB) CompleteUpload(id string) error {
	_, err := db.conn.Exec(
		`
		UPDATE uploads
		SET status = ?
		WHERE id = ?
		
		`,
		UploadComplete,
		id,
	)
	return err
}
