package encoder

import (
	"github.com/mrod502/encmsg/util"
)

// New instantiates an encoder with encrypter enc and serializer srl
func New(enc Encrypter, srl Serializer) *Encoder {
	return &Encoder{
		Encrypter: enc,
		srl:       srl,
	}
}

// Encoder is a wrapper for an Encrypter that handles serialization and deserialization of data,
// and then chunks the serialized bytes, then encrypts them and then stores the encrypted data
// in a Message object
type Encoder struct {
	Encrypter

	srl Serializer
}

// Encode serializes `val` using the defined serializer,
// then encrypts the data with the provided encrypter
func (e *Encoder) Encode(val interface{}) ([]byte, error) {
	b, err := e.srl(val)
	if err != nil {
		return nil, err
	}

	chunks := util.ToChunks(b, e.MaxChunkSize())
	encChunks := make([][]byte, 0, len(chunks))

	for _, chunk := range chunks {
		res, err := e.Encrypt(chunk)
		if err != nil {
			return nil, err
		}
		encChunks = append(encChunks, res)
	}

	out, err := e.srl(util.Message{Data: encChunks, Nonce: e.Nonce()})
	if err == nil {
		e.UpdateNonce(out)
	}
	return out, err

}
