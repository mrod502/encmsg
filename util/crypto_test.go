package util

import (
	"fmt"
	"testing"
)

func fooChunk(b []byte) (data [][]byte, err error) {

	chunkSize := 3
	data = make([][]byte, 0, len(b)/chunkSize)

	for i := 0; i < len(b); i += chunkSize {
		res := b[i:min(i+chunkSize, len(b))]

		data = append(data, res)

	}
	return
}

func TestChunky(t *testing.T) {

	var foo = []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xf, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18}

	b, _ := fooChunk(foo)
	fmt.Printf("%+v\n", b)

}
