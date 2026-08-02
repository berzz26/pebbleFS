# ADR-0002: Storage Layout

## Status

Accepted

## Context

PebbleFS needs a location on each registered storage device to persist chunk data.

One option was writing directly to raw block devices.

Another option was storing chunks inside an existing filesystem.

## Decision

PebbleFS will create a hidden directory on every registered storage device.

```
.pebble/
    disk.json
    chunks/
```

`disk.json` stores persistent metadata about the storage device.

The underlying filesystem (ext4, NTFS, exFAT, etc.) remains responsible for low-level storage management.

## Why

- Simpler implementation.
- Cross-platform.
- Reuses decades of filesystem reliability.
- Easier debugging.
- Easier migration between machines.

## Consequences

PebbleFS becomes a distributed storage layer rather than replacing the local filesystem.