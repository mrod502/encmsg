package encoder

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/mrod502/encmsg/util"
)

func TestRsaEncrypter(t *testing.T) {
	key, _ := util.GenerateRsaKey(4096)
	enc := NewRsaEncrypter(&key.PublicKey)

	theMessage := []byte(`hello this is a message that I want to encrypt, please encrypt me`)

	b, err := enc.Encrypt(theMessage)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(b))
}
