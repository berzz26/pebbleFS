package metadata

func (db *DB) AddChunkToFile(
	fileID int64,
	chunkID string,
	index int,
) error {

	_, err := db.conn.Exec(
		`
		INSERT INTO file_chunks(
			file_id,
			chunk_id,
			chunk_index
		)
		VALUES(?,?,?)
		`,
		fileID,
		chunkID,
		index,
	)

	return err
}