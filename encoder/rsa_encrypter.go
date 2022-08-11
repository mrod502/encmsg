package encoder

import (
	crand "crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/binary"
	"math/rand"
	"sync"
	"time"
)

type RsaEncrypter struct {
	*sync.RWMutex
	pub *rsa.PublicKey

	nonce []byte
}

func NewRsaEncrypter(pub *rsa.PublicKey) *RsaEncrypter {
	r:= &RsaEncrypter{
		RWMutex: &sync.RWMutex{},
		pub:     pub,
	}
	r.nonce = r.newNonce()
	return r
}

func (e *RsaEncrypter) Encrypt(chunk []byte) ([]byte, error) {
	e.RLock()
	defer e.RUnlock()
	return rsa.EncryptOAEP(
		sha256.New(),
		crand.Reader,
		e.pub,
		append(e.nonce, chunk...),
		nil,
	)
}

func (e RsaEncrypter) Nonce() []byte {
	return append([]byte{},e.nonce...)
}

func (e RsaEncrypter) MaxChunkSize() int {
	return e.pub.Size() - (66 + len(e.nonce))
}

func (e *RsaEncrypter) UpdateNonce(data []byte) {
	e.Lock()
	defer e.Unlock()
	e.nonce = e.newNonce()
}

func (e RsaEncrypter) newNonce() []byte {
	rand := uint64(rand.Int63() * time.Now().UnixNano())

	bs := make([]byte, 8)
	binary.LittleEndian.PutUint64(bs, rand)
	return bs
}
