# ADR-0004: Directory Sharding

## Status

Accepted

## Context

Storing every chunk inside a single directory would eventually create millions of files in one location.

Large directories become increasingly inefficient to manage and search.

## Decision

Chunk paths are derived from the hash.

Example:

```
Hash

7d4e9c4d0f...

↓

.pebble/chunks/7d/4e/7d4e9c4d0f...
```

The first two bytes form the first directory.

The next two bytes form the second directory.

The full hash becomes the filename.

## Why

- Prevents huge directories.
- Keeps filesystem metadata manageable.
- Deterministic path calculation.
- No database lookup required to locate a chunk.

## Consequences

Chunk location can always be recomputed from its hash.

No additional path indexing is required.