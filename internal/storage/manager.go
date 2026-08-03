package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

type Manager struct {
	disks map[string]*Disk
}

func NewManager() *Manager {
	return &Manager{
		disks: make(map[string]*Disk),
	}
}

func (m *Manager) AddDisk(path string) (*Disk, error) {

	meta, err := loadMetadata(path)

	if err != nil {

		// treat only errors that represent missing .json file
		if !os.IsNotExist(err) {
			return nil, err
		}

		pebbleDir := filepath.Join(path, ".pebble")
		chunkDir := filepath.Join(pebbleDir, "chunks")

		if err := os.MkdirAll(chunkDir, 0755); err != nil {
			return nil, err
		}

		meta = createDiskMetadata(path)

		if err := saveMetadata(path, meta); err != nil {
			return nil, err
		}

	}

	total, free, err := GetDiskStats(path)
	if err != nil {
		return nil, err
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
		return nil, fmt.Errorf("disk %s already added", meta.ID)
	}
	m.disks[disk.ID] = disk

	return disk, nil
}

// why this? store the id of each mount on the disk itself, eaach disk gets disk.json where meta of the disk is stored
// this json file will be the first thinkg pebbleFs reads and knows info abt disk
func createDiskMetadata(path string) *DiskMetadata {

	return &DiskMetadata{
		ID:         uuid.New().String(),
		Version:    1,
		CreatedAt:  time.Now(),
		Filesystem: detectFilesystem(path),
		Label:      filepath.Base(path),
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

func (m *Manager) Disks() []*Disk {
	// returns a pool of disk from the private map.
	disks := make([]*Disk, 0, len(m.disks))

	for _, disk := range m.disks {
		disks = append(disks, disk)
	}

	return disks
}

func (m *Manager) GetDisk(id string) (*Disk, bool) {
	disk, ok := m.disks[id]
	return disk, ok
}

// after a write, update the available space metadata of the disk
func (m *Manager) ReserveSpace(diskID string, bytes uint64) error {

	disk, ok := m.disks[diskID]
	if !ok {
		return errors.New("disk not found")
	}

	if disk.FreeSpace < bytes {
		return errors.New("not enough space")
	}

	disk.FreeSpace -= bytes

	return nil
}

func (m *Manager) Count() int {
	return len(m.disks)
}

func (m *Manager) TotalCapacity() uint64 {

	var total uint64

	for _, disk := range m.disks {
		total += disk.TotalSpace
	}

	return total
}
func (m *Manager) TotalFreeSpace() uint64 {

	var total uint64

	for _, disk := range m.disks {
		total += disk.FreeSpace
	}

	return total
}
func (m *Manager) RemoveDisk(id string) {

	delete(m.disks, id)
}


//will implement this later.
func detectFilesystem(path string) string {
	return "unkown"
}

func (m *Manager) ReleaseSpace(
	diskID string,
	bytes uint64,
) error {

	disk, ok := m.disks[diskID]
	if !ok {
		return errors.New("disk not found")
	}

	disk.FreeSpace += bytes

	return nil
}
