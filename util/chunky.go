package util

type ChunkReader interface {
	Read([][]byte) ([]byte, error)
}

type ChunkWriter interface {
	Write([]byte) ([][]byte, error)
}

type RsaChunkReader struct {
}
