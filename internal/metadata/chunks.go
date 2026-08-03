package metadata


func (db *DB) UpsertChunk(chunk Chunk) (bool, error) {

	err := db.InsertChunk(chunk)
	if err == nil {
		return true, nil
	}

	err = db.IncrementReference(chunk.ID)
	if err != nil {
		return false, err
	}

	return false, nil
}

func (db *DB) ChunkExists(id string) (bool, error) {

	var exists int

	err := db.conn.QueryRow(
		`
		SELECT EXISTS(
			SELECT 1
			FROM chunks
			WHERE id = ?
		)
		`,
		id,
	).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists == 1, nil
}

func (db *DB) InsertChunk(chunk Chunk) error {

	_, err := db.conn.Exec(
		`
		INSERT INTO chunks(
			id,
			disk_id,
			path,
			size,
			reference_count
		)
		VALUES(?,?,?,?,?)
		`,
		chunk.ID,
		chunk.DiskID,
		chunk.Path,
		chunk.Size,
		1,
	)

	return err
}

func (db *DB) IncrementReference(id string) error {

	_, err := db.conn.Exec(
		`
		UPDATE chunks
		SET reference_count = reference_count + 1
		WHERE id = ?
		`,
		id,
	)

	return err
}
func (db *DB) DecrementReference(id string) (int64, error) {

	_, err := db.conn.Exec(`
		UPDATE chunks
		SET reference_count = reference_count - 1
		WHERE id = ?
	`, id)

	if err != nil {
		return 0, err
	}

	var refs int64

	err = db.conn.QueryRow(`
		SELECT reference_count
		FROM chunks
		WHERE id = ?
	`, id).Scan(&refs)

	if err != nil {
		return 0, err
	}

	return refs, nil
}

func (db *DB) DeleteChunk(id string) error {

	_, err := db.conn.Exec(`
		DELETE FROM chunks
		WHERE id = ?
	`, id)

	return err
}