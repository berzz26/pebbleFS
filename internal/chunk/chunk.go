package chunk

type Chunk struct {
	//why id and index both? coz later, this chunk id will be useful for replication, checksums, dedupe, and content-addressed storage if reqiuired
	ID string
	Index int
	
	Data  []byte
	Size  int64
}

