package checksum

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"strings"
	"testing"
)

func TestReadAndVerify(t *testing.T) {
	input := []byte("hello")
	checksum, _ := hex.DecodeString("2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824")
	reader := NewReader(bufio.NewReader(bytes.NewBuffer(append(input, checksum...))))
	dataRead, _ := ioutil.ReadAll(reader)
	if !reader.Verify() {
		t.Fatalf("could not pass data verification")
	}
	if !bytes.Equal(dataRead, input) {
		t.Fatalf("could not read correct data")
	}
}

func TestReadWithIncorrectHash(t *testing.T) {
	input := []byte("a long msg")
	checksum := make([]byte, sha256.Size)
	reader := NewReader(bufio.NewReader(bytes.NewBuffer(append(input, checksum...))))
	dataRead, _ := ioutil.ReadAll(reader)
	if reader.Verify() {
		t.Fatalf("should report as failure")
	}
	if !bytes.Equal(dataRead, input) {
		t.Fatalf("could not read correct data")
	}
}

func TestReadWithInsufficientData(t *testing.T) {
	input := []byte("a long msg")
	reader := NewReader(bufio.NewReader(bytes.NewBuffer(input)))
	dataRead, readError := ioutil.ReadAll(reader)
	if len(dataRead) != 0 {
		t.Fatalf("should not read anything")
	}
	if readError == nil {
		t.Fatalf("should trigger a read error")
	}
	if !strings.Contains(readError.Error(), "insufficient data") {
		t.Fatalf("should complain about insufficient data")
	}
	println(readError.Error())
	if reader.Verify() {
		t.Fatalf("should report as failure")
	}
}
