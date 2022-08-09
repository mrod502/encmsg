package encmsg

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
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
	return decoder.New(decrypter, json.Unmarshal, json.Unmarshal)
}

func Encoder(key *rsa.PrivateKey) *encoder.Encoder {
	encrypter := encoder.NewRsaEncrypter(&key.PublicKey)
	return encoder.New(encrypter, json.Marshal, json.Marshal)
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
			B: []byte("this is another sentence"),
		},
	}

	dec := Decoder(key)
	enc := Encoder(key)
	fmt.Println("max chunk size", enc.MaxChunkSize())

	b, err := enc.Encode(val)
	if err != nil {
		t.Fatal(err)
	}
	var decStruct TestStruct
	fmt.Println("ENCODED:", string(b))
	err = dec.Decode(b, &decStruct)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n%+v\n", val, decStruct)

}
