# ADR-0001: PebbleFS Vision

## Status

Accepted

## Context

The goal of PebbleFS is not to build another traditional filesystem.

Instead, PebbleFS is designed as a distributed storage engine capable of aggregating arbitrary storage devices (USBs, HDDs, SSDs, etc.) into a single logical storage pool.

A filesystem interface (FUSE) will be added later as a presentation layer.

## Decision

PebbleFS will be built in layers.

Storage Devices
↓

Chunk Storage
↓

Metadata
↓

Placement
↓

Replication
↓

Filesystem Interface

Each layer has a single responsibility and should remain independent.

## Consequences

- Easier to extend.
- Networking can be added later.
- FUSE remains optional.
- Components can be tested independently.