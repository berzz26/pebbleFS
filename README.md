PebbleFS is my attempt to build a distributed storage engine completely from scratch.

The goal is to support almost any storage device that Linux can **mount**—USB drives, SSDs, HDDs, NVMe drives, or storage attached to another machine—and combine them into a single self-managing storage pool.

As the project evolves, PebbleFS should automatically discover new storage, distribute data across devices, tolerate failures, replicate important data, recover from lost disks, and continue growing as additional machines join the cluster.