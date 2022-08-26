package decoder

import (
	"github.com/mrod502/encmsg/util"
)

// Decoder takes a stream of bytes (a `Message` instance), decrypts
// using the provided decrypter, and then deserializes the data to an
// instance of the original object that was serialized
type Decoder struct {
	Decrypter
	deser Deserializer
}

// New initializes a new Decoder
func New(dec Decrypter, deser Deserializer) *Decoder {
	return &Decoder{
		Decrypter: dec,
		deser:     deser,
	}
}

// Decode decrypts `b` and deserializes the data to `v`
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

	return d.deser(out, v)
}
