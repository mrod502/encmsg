package encoder

type Encrypter interface {
	Encrypt([]byte) ([]byte, error)
	MaxChunkSize() int // max chunk size, in bytes
	Nonce() []byte
	UpdateNonce([]byte)
}

type Serializer func(interface{}) ([]byte, error)
