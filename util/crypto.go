package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
)

func ParseRsaPublicKeyFromPem(pubPEM []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pubPEM))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	switch pub := pub.(type) {
	case *rsa.PublicKey:
		return pub, nil
	default:
		break // fall through
	}
	return nil, errors.New("key type is not RSA")
}

func ExportRsaPublicKey(pubkey *rsa.PublicKey) (*pem.Block, error) {
	pubkey_bytes, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		return nil, err
	}

	return &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubkey_bytes,
	}, nil
}

func ParseRsaPrivateKeyFromPem(privPEM []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privPEM))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return priv, nil
}

func GenerateRsaKey(size int) (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, size)

}

func LoadPubKey(path string) (*rsa.PublicKey, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return ParseRsaPublicKeyFromPem(b)
}

func LoadPrivKey(path string) (*rsa.PrivateKey, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return ParseRsaPrivateKeyFromPem(b)
}

func ToChunks(b []byte, chunkSize int) (data [][]byte) {
	data = make([][]byte, 0, len(b)/chunkSize)
	for i := 0; i < len(b); i += chunkSize {
		data = append(data, b[i:min(i+chunkSize, len(b))])
	}
	return
}

func MaxChunkSize(key *rsa.PublicKey) int {
	return key.Size() - 42
}
