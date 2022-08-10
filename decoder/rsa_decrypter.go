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
	b, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, d.priv, data, nil)
	if err != nil {
		return nil, err
	}
	return b[len(nonce):], nil
}
