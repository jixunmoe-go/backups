package checksum

import (
	"crypto/sha256"
	"fmt"
	"hash"
	"io"
)

type Writer struct {
	writer io.Writer
	hash   hash.Hash
}

func NewWriter(writer io.Writer) *Writer {
	return &Writer{
		writer: writer,
		hash:   sha256.New(),
	}
}

func (w *Writer) Write(p []byte) (int, error) {
	w.hash.Write(p)
	return w.writer.Write(p)
}

func (w *Writer) WriteChecksum() (int, error) {
	n, err := w.writer.Write(w.hash.Sum(nil))
	if err != nil {
		return 0, err
	}
	if w.hash.Size() != n {
		return n, fmt.Errorf("expected to write %d bytes checksum, written %d bytes instead", w.hash.Size(), n)
	}
	return n, nil
}
