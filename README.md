# encmsg

Support for chunked encryption/decryption of data, with builtin support for RSA Private/Public key pairs.

## Usage

```go
import (
    "encoding/json"

    "github.com/mrod502/encmsg/encoder"
    "github.com/mrod502/encmsg/decoder"
    "github.com/mrod502/encmsg/util"
)

func Decoder(key *rsa.PrivateKey) *decoder.Decoder {
	decrypter := decoder.NewRsaDecrypter(key)
	return decoder.New(decrypter, json.Unmarshal)
}

func Encoder(key *rsa.PublicKey) *encoder.Encoder {
	encrypter := encoder.NewRsaEncrypter(key)
	return encoder.New(encrypter, json.Marshal)
}

type S struct {
    Field1 string
    Field2 uint64
    Field3 []byte
}

func main() {
    key, _ := util.GenerateRsaKey(4096)

    val := S{
        Field1: "themeaningoflife",
        Field2: 69,
        Field3: []byte{0xff, 0xff, 0xee},
    }
    // the sender
    enc := Encoder(&key.PublicKey)

    //the receiver
    dec := Decoder(key)

    b, _ := enc.Encode(val)

    var decStruct TestStruct
    err = dec.Decode(b, &decStruct)
    if err != nil {
        panic(err)
    }
}
```
