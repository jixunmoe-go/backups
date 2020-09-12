package crypto

import (
	"bytes"
	"testing"
)

func TestDeriveEncryptionKey(t *testing.T) {
	priv1, pub1, _ := GenKeyPair()
	priv2, pub2, _ := GenKeyPair()

	if bytes.Equal(priv1, priv2) {
		t.Fatalf("should not generate the same private key (rng issue?)")
	}

	key1, _ := DeriveEncryptionKey(priv1, pub2)
	key2, _ := DeriveEncryptionKey(priv2, pub1)

	if !bytes.Equal(key1, key2) {
		t.Fatalf("could not derive correct key")
	}
}
