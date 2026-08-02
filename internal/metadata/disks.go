package metadata

import "github.com/berzz/pebbleFS/internal/storage"

func (db *DB) RegisterDisk(disk storage.Disk) error {

	_, err := db.conn.Exec(
		`
		INSERT INTO disks(
			id,
			mount_path,
			total_space,
			free_space,
			healthy
		)
		VALUES(?,?,?,?,?)
		`,
		disk.ID,
		disk.MountPath,
		disk.TotalSpace,
		disk.FreeSpace,
		disk.Healthy,
	)

	return err
}

func (db *DB) GetDisks() ([]storage.Disk, error) {

	rows, err := db.conn.Query(
		`
		SELECT
			id,
			mount_path,
			total_space,
			free_space,
			healthy
		FROM disks
		`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var disks []storage.Disk

	for rows.Next() {

		var d storage.Disk

		err := rows.Scan(
			&d.ID,
			&d.MountPath,
			&d.TotalSpace,
			&d.FreeSpace,
			&d.Healthy,
		)

		if err != nil {
			return nil, err
		}

		disks = append(disks, d)
	}

	return disks, nil
}
