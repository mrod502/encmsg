package encmsg

import (
	"crypto/rsa"
	"encoding/json"
	"strings"
	"testing"

	"github.com/mrod502/encmsg/decoder"
	"github.com/mrod502/encmsg/encoder"
	"github.com/mrod502/encmsg/util"
)

type Mini struct {
	A uint32
	B []byte
}
type TestStruct struct {
	SomeField      uint64
	SomeOtherField string
	SomeStruct     Mini
}

func Decoder(key *rsa.PrivateKey) *decoder.Decoder {
	decrypter := decoder.NewRsaDecrypter(key)
	return decoder.New(decrypter, json.Unmarshal)
}

func Encoder(key *rsa.PublicKey) *encoder.Encoder {
	encrypter := encoder.NewRsaEncrypter(key)
	return encoder.New(encrypter, json.Marshal)
}

func TestEncodeDecode(t *testing.T) {
	key, err := util.GenerateRsaKey(4096)
	if err != nil {
		t.Fatal(err)
	}

	var val = TestStruct{
		SomeField:      123456789,
		SomeOtherField: "hello world it's me",
		SomeStruct: Mini{
			A: 123456,
			B: []byte(strings.Repeat("this is another sentence", 1999)),
		},
	}

	dec := Decoder(key)
	enc := Encoder(&key.PublicKey)

	b, err := enc.Encode(val)
	if err != nil {
		t.Fatal(err)
	}
	var decStruct TestStruct
	err = dec.Decode(b, &decStruct)
	if err != nil {
		t.Fatal(err)
	}
}
