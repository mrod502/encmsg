package encoder

import (
	"github.com/mrod502/encmsg/util"
)

func New(enc Encrypter, srl Serializer, encoder Serializer) *Encoder {
	return &Encoder{
		Encrypter: enc,
		srl:       srl,
		encode:    encoder,
	}
}

type Encoder struct {
	Encrypter

	srl    Serializer
	encode Serializer // encodes the message
}

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

	out, err := e.encode(util.Message{Data: encChunks, Nonce: e.Nonce()})
	if err == nil {
		e.UpdateNonce(out)
	}
	return out, err

}
