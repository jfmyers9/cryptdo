package cryptdo_test

import (
	"testing"

	"github.com/golang/protobuf/proto"

	"code.xoeb.us/cryptdo"
	"code.xoeb.us/cryptdo/cryptdopb"
)

func TestCurrentCrypto(t *testing.T) {
	passphrase := "hunter2"
	plaintext := []byte("something secret")

	ciphertext, err := cryptdo.Encrypt(plaintext, passphrase)
	if err != nil {
		t.Error("got error while encrypting:", err)
	}

	message := &cryptdopb.Message{}
	if err := proto.Unmarshal(ciphertext, message); err != nil {
		t.Error("got error while unmarshaling", err)
	}

	if version := message.GetVersion(); version != 1 {
		t.Errorf("version was incorrect, got: %d, want: %d", version, 1)
	}
}
