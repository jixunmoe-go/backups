package checksum

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestWriterChecksumOutput(t *testing.T) {
	input := []byte("hello")
	checksum, _ := hex.DecodeString("2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824")
	var buf bytes.Buffer
	writer := NewWriter(&buf)

	_, _ = writer.Write(input)
	_, _ = writer.WriteChecksum()

	expected := append(input, checksum...)
	actual := buf.Bytes()

	if !bytes.Equal(expected, actual) {
		t.Fatalf("did not write checksum after content")
	}
}
