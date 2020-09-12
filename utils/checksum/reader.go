package checksum

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"hash"
	"io"
)

type Reader struct {
	reader io.Reader
	hash   hash.Hash
	buffer []byte
	tmp    []byte
}

func NewReader(reader io.Reader) *Reader {
	return &Reader{
		reader: reader,
		hash:   sha256.New(),
		buffer: nil,
		tmp:    nil,
	}
}

func (r *Reader) Read(p []byte) (n int, err error) {
	bufSize := len(r.buffer)
	readSize := sha256.Size - bufSize + len(p)
	if len(r.tmp) != readSize {
		r.tmp = make([]byte, readSize)
	}
	n, err = r.reader.Read(r.tmp)

	dataAvailable := append(r.buffer, r.tmp[0:n]...)

	// insufficient data.
	if len(dataAvailable) < sha256.Size {
		return 0, fmt.Errorf("insufficient data, need %d bytes, got %d bytes", sha256.Size, len(dataAvailable))
	}

	actualRead := len(dataAvailable) - sha256.Size
	r.buffer = dataAvailable[actualRead:]
	dataRead := dataAvailable[0:actualRead]
	copy(p, dataRead)

	r.hash.Write(dataRead)
	return actualRead, err
}

func (r *Reader) Verify() bool {
	return len(r.buffer) == r.hash.Size() && bytes.Equal(r.hash.Sum(nil), r.buffer)
}
