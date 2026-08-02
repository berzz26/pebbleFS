package chunk

import "path/filepath"


// this function shards the /chunk directory to smaller edirectories so that search operation becomes less expensive when searching
// why hashed prefixes? sha256 is 64 hexadecimal chars. with that, we get exactly 256 top level dirs and hashes help distribute large num of files uniformly
func ChunkPath(root string, hash string) string {
	level1 := hash[:2]
	level2 := hash[2:4]

	return filepath.Join(
		root,
		level1,
		level2,
		hash,
	)
}