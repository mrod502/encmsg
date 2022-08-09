package decoder

import (
	"github.com/mrod502/encmsg/util"
)

type Decoder struct {
	Decrypter
	deser       Deserializer
	structDeser Deserializer
}

func New(dec Decrypter, deser Deserializer, sdeser Deserializer) *Decoder {
	return &Decoder{
		Decrypter:   dec,
		deser:       deser,
		structDeser: sdeser,
	}
}

func (d *Decoder) Decode(b []byte, v interface{}) error {
	var msg util.Message
	if err := d.deser(b, &msg); err != nil {
		return err
	}
	out := make([]byte, 0)

	for _, chunk := range msg.Data {
		decrypted, err := d.Decrypt(chunk, msg.Nonce)
		if err != nil {
			return err
		}
		out = append(out, decrypted...)
	}

	return d.structDeser(out, v)
}
