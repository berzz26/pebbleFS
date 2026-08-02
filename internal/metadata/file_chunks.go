package metadata

type FileChunkEntry struct {
	ChunkID string
	Index   int
}

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

func (db *DB) GetFileChunks(fileID int64) ([]FileChunkEntry, error) {

	rows, err := db.conn.Query(
		`
		SELECT chunk_id, chunk_index
		FROM file_chunks
		WHERE file_id = ?
		ORDER BY chunk_index
		`,
		fileID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var chunks []FileChunkEntry

	for rows.Next() {

		var c FileChunkEntry

		if err := rows.Scan(&c.ChunkID, &c.Index); err != nil {
			return nil, err
		}

		chunks = append(chunks, c)
	}

	return chunks, nil
}