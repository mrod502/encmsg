package encoder

import (
	crand "crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/json"
	"math/rand"
	"sync"
)

type RsaEncrypter struct {
	*sync.RWMutex
	pub *rsa.PublicKey

	nonce []byte
}

func NewRsaEncrypter(pub *rsa.PublicKey) *RsaEncrypter {
	num := rand.Int63n(1 << 62)
	b, _ := json.Marshal(num)
	nonce := sha256.Sum256(b)
	return &RsaEncrypter{
		RWMutex: &sync.RWMutex{},
		pub:     pub,
		nonce:   nonce[:],
	}
}

func (e *RsaEncrypter) Encrypt(chunk []byte) ([]byte, error) {
	e.RLock()
	defer e.RUnlock()

	return rsa.EncryptOAEP(
		sha256.New(),
		crand.Reader,
		e.pub,
		chunk,
		[]byte("msg"),
	)
}

func (e RsaEncrypter) Nonce() []byte {
	return e.nonce
}

func (e RsaEncrypter) MaxChunkSize() int {
	return e.pub.Size() - 66
}

func (e *RsaEncrypter) UpdateNonce(data []byte) {
	e.Lock()
	defer e.Unlock()
	b := sha256.Sum256(append(e.nonce, data...))
	rand, _ := json.Marshal(rand.Int63n(1 << 62))

	b = sha256.Sum256(append(b[:], rand...))
	e.nonce = b[:]
}
