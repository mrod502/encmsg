package decoder

type Deserializer func([]byte, interface{}) error

type Decrypter interface {
	Decrypt(b, nonce []byte) ([]byte, error)
}
