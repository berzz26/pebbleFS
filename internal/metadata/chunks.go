package metadata


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