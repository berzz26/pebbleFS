# Changelog

## v0.1.0

Initial public release.

### Added

- Multi-device storage manager
- Persistent disk identities
- SHA-256 content-addressed chunking
- Directory sharding
- Chunk writer
- Chunk reader
- Upload pipeline
- Download pipeline
- File metadata
- Chunk metadata
- File-to-chunk mapping
- Upload tracking
- Placement strategy
- File listing
- File information
- File deletion
- Disk registration
- Disk listing
- Cobra CLI

### Known Limitations

- Single-node only
- SQLite write contention under concurrent writes
- No replication
- No integrity verification
- No crash recovery
- No FUSE interface