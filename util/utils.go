package util

type number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func min[T number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

type Message struct {
	Data  [][]byte
	Nonce []byte
}
