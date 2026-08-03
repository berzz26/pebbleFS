# PebbleFS

PebbleFS is an experimental distributed storage engine written in Go.

Instead of relying on a single storage device, PebbleFS aggregates multiple storage devices (USB drives, HDDs, SSDs, flash drives, etc.) into a single logical storage pool.

Files are chunked, content-addressed using SHA-256, distributed across registered disks, and reconstructed transparently during reads.

> PebbleFS is currently a storage engine. A FUSE filesystem interface and distributed clustering are planned for future releases.

---

## Features

- Multi-disk storage pooling
- Content-addressed storage (SHA-256)
- Automatic chunk deduplication
- Persistent disk registration
- Chunk directory sharding
- File upload
- File download
- File deletion
- File listing
- File information
- Storage statistics
- Service-oriented architecture

---

## Architecture

```
             CLI

              │

        Service Layer

      ┌───────┼────────┐

 Metadata   Storage   Chunk

(SQLite)   Manager    Engine
```

Each layer owns one responsibility.

---

## Storage Layout

Each registered disk contains:

```
.pebble/

    disk.json

    chunks/

        7d/

            4e/

                7d4e9c...
```

Chunks are stored using deterministic directory sharding.

---

## Installation

```bash
git clone ...

go build -o pebble ./cmd/pebble
```

---

## Usage

Register a disk

```bash
./pebble disk add /mnt/usb
```

List disks

```bash
./pebble disk list
```

Upload

```bash
./pebble put movie.mp4
```

Download

```bash
./pebble get movie.mp4
```

Download elsewhere

```bash
./pebble get movie.mp4 recovered.mp4
```

List files

```bash
./pebble list
```

File info

```bash
./pebble info movie.mp4
```

Delete

```bash
./pebble rm movie.mp4
```

---

## Current Status

PebbleFS currently supports:

- Single node
- Multiple storage devices
- Persistent metadata
- Deduplicated chunk storage

---

## Roadmap

### v0.2

- Verification
- Crash recovery
- WAL mode
- Garbage collection
- Resumable uploads

### v0.3

- Replication
- Background healing
- Automatic rebalancing

### v1.0

- Distributed cluster
- Networked nodes
- FUSE filesystem
- Full storage pool abstraction