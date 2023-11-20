package mangler

import (
	"bytes"
	"fmt"
	"math/rand"
	"testing"
)

func TestMangler(t *testing.T) {
	mangler := NewMangler()
	defer mangler.Close()
}

func TestEncryptAndDecrypt(t *testing.T) {
	key := []byte(fmt.Sprintf("%016x", rand.Uint64()))
	message := []byte(`hello world!`)
	encrypted := func() []byte {
		mangler := NewMangler()
		for _, c := range key {
			mangler.Encrypt(c)
		}
		buffer := bytes.Buffer{}
		for _, c := range message {
			buffer.WriteByte(mangler.Encrypt(c))
		}
		return buffer.Bytes()
	}()
	decypted := func() []byte {
		mangler := NewMangler()
		for _, c := range key {
			mangler.Encrypt(c)
		}
		buffer := bytes.Buffer{}
		for _, c := range encrypted {
			buffer.WriteByte(mangler.Decrypt(c))
		}
		return buffer.Bytes()
	}()
	if !bytes.Equal(decypted, message) {
		t.Fatalf("error: wanted '%s'; got '%s'\n", message, decypted)
	}
}
