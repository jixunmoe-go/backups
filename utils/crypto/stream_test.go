package crypto

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"strings"
	"testing"
)

func readFile(t *testing.T, name string) []byte {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		t.Fatalf("could not get file: %s", err)
	}
	return data
}

func readB64(t *testing.T, name string) []byte {
	data := readFile(t, name)
	decoded, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		t.Fatalf("could not parse file: %s", err)
	}

	return decoded
}

func TestEncryptAndDecryptStream(t *testing.T) {
	var pubKey PublicKey = readB64(t, "testdata/pubkey.txt")
	var privKey PrivateKey = readB64(t, "testdata/privkey.txt")

	var cryptedBuf bytes.Buffer
	err := EncryptStream(bytes.NewBufferString("hello!"), &cryptedBuf, pubKey)
	if err != nil {
		t.Fatalf("could not encrypt our virtual buffer: %s", err)
	}

	var decryptedBuf bytes.Buffer
	err = DecryptStream(&cryptedBuf, &decryptedBuf, privKey)

	if err != nil {
		t.Fatalf("decrypt fail: %s", err)
	}

	if !bytes.Equal(decryptedBuf.Bytes(), []byte("hello!")) {
		t.Fatalf("decrypted data differ")
	}
}

func TestDecryptDataWithGoodHash(t *testing.T) {
	privKey := readB64(t, "testdata/privkey.txt")
	encrypted := readFile(t, "testdata/encrypted_hash_ok.bin")

	var decryptedBuf bytes.Buffer
	err := DecryptStream(bytes.NewBuffer(encrypted), &decryptedBuf, privKey)
	if err != nil {
		t.Fatalf("decrypt fail: %s", err)
	}

	if !bytes.Equal(decryptedBuf.Bytes(), []byte("hello!\n")) {
		t.Fatalf("decrypted data differ")
	}
}

func TestDecryptDataWithBadHash(t *testing.T) {
	privKey := readB64(t, "testdata/privkey.txt")
	encrypted := readFile(t, "testdata/encrypted_hash_bad.bin")

	var decryptedBuf bytes.Buffer
	err := DecryptStream(bytes.NewBuffer(encrypted), &decryptedBuf, privKey)
	if err == nil {
		t.Fatalf("decrypt should report fail but it didn't")
	}

	if !strings.Contains(err.Error(), "could not verify") {
		t.Fatalf("it should report verifycation error")
	}
}
