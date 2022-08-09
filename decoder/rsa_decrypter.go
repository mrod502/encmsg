package decoder

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

type RsaDecrypter struct {
	priv *rsa.PrivateKey
}

func NewRsaDecrypter(priv *rsa.PrivateKey) *RsaDecrypter {
	return &RsaDecrypter{
		priv: priv,
	}
}

func (d RsaDecrypter) Decrypt(data, nonce []byte) ([]byte, error) {
	return rsa.DecryptOAEP(sha256.New(), rand.Reader, d.priv, data, []byte("msg"))
}
