package storage

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"os"
	"path/filepath"
)

type Manager struct {
	disks map[string]*Disk
}

func NewManager() *Manager {
	return &Manager{
		disks: make(map[string]*Disk),
	}
}

func (m *Manager) AddDisk(path string) error {

	meta, err := loadMetadata(path)

	if err != nil {

		// treat only errors that represent missing .json file
		if !os.IsNotExist(err) {
			return err
		}

		pebbleDir := filepath.Join(path, ".pebble")
		chunkDir := filepath.Join(pebbleDir, "chunks")

		if err := os.MkdirAll(chunkDir, 0755); err != nil {
			return err
		}

		meta = createDiskMetadata()

		if err := saveMetadata(path, meta); err != nil {
			return err
		}

	}

	total, free, err := GetDiskStats(path)
	if err != nil {
		return err
	}

	disk := &Disk{
		ID:         meta.ID,
		MountPath:  path,
		Healthy:    true,
		TotalSpace: total,
		FreeSpace:  free,
	}
	//make sure that overwriting the map entry doesnt succeed
	if _, exists := m.disks[meta.ID]; exists {
    return fmt.Errorf("disk %s already added", meta.ID)
}
	m.disks[disk.ID] = disk

	fmt.Printf("Added disk %s at %s\n", disk.ID, path)

	return nil
}

// why this? store the id of each mount on the disk itself, eaach disk gets disk.json where meta of the disk is stored
// this json file will be the first thinkg pebbleFs reads and knows info abt disk
func createDiskMetadata() *DiskMetadata {

	return &DiskMetadata{
		ID:      uuid.New().String(),
		Version: 1,
	}


}

func saveMetadata(path string, meta *DiskMetadata) error {

	file, err := os.Create(filepath.Join(path, ".pebble", "disk.json"))
	if err != nil {
		return err
	}

	defer file.Close()

	return json.NewEncoder(file).Encode(meta)
}

func loadMetadata(path string) (*DiskMetadata, error) {

	file, err := os.Open(filepath.Join(path, ".pebble", "disk.json"))
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var meta DiskMetadata

	if err := json.NewDecoder(file).Decode(&meta); err != nil {
		return nil, err
	}

	return &meta, nil
}
