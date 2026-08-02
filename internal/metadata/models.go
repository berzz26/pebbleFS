package metadata

func (d *DB) initialize() error {

	query := `
	CREATE TABLE IF NOT EXISTS files(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		filename TEXT NOT NULL,
		size INTEGER NOT NULL
	);

	CREATE TABLE IF NOT EXISTS chunks(
		id TEXT PRIMARY KEY,
		file_id INTEGER,
		disk_id TEXT,
		path TEXT,
		chunk_index INTEGER,
		size INTEGER,
		checksum TEXT
	);
	`
	_,err := d.conn.Exec(query)

	return  err
}
