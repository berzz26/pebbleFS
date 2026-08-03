# PebbleFS Architecture

PebbleFS is organized into independent layers.

```
CLI

↓

Service Layer

↓

Metadata

Storage Manager

Placement

Chunk Engine

↓

Physical Storage
```

## Chunk Engine

Responsible for:

- splitting files
- hashing
- reading
- writing
- deleting chunks

## Metadata

Responsible for:

- logical files
- physical chunks
- file mappings
- uploads
- registered disks

## Storage Manager

Responsible for:

- registered disks
- free space
- health
- capacity

## Placement

Responsible for deciding where chunks should be stored.

## Service Layer

Coordinates every subsystem.

It contains the application's business logic.