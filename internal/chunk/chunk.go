package chunk

type Chunk struct {
	//why id and index both? coz later, this chunk id will be useful for replication, checksums, dedupe, and content-addressed storage if reqiuired
	// initially, this was a uuid, but later we shifted to a sha 256 string to reduce overhead and save space
	ID string
	Index int
	
	Data  []byte
	Size  int64
}

