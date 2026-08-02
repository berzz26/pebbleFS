package storage

type Disk struct {
	ID         string
	MountPath  string
	TotalSpace uint64
	FreeSpace  uint64
	Healthy    bool
}
