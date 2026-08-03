# Storage Model

PebbleFS separates physical storage from logical files.

## Physical Layer

```
Chunks

↓

Disk

↓

Path

↓

Reference Count
```

Physical chunks exist exactly once.

---

## Logical Layer

```
File

↓

Chunk References

↓

Ordered Mapping
```

Files do not own chunk data.

Instead, they reference physical chunks.

This enables:

- deduplication
- snapshots
- cloning
- copy-on-write
- garbage collection

---

## Chunk IDs

Chunk IDs are SHA-256 hashes of chunk contents.

Properties:

- deterministic
- content-addressed
- globally unique (practically)
- integrity verification
- deduplication