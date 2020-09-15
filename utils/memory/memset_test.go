package memory

import (
	"crypto/rand"
	"testing"
)

const arraySize = 32
const boundaryBufSize = zeroBufSize * 2
const largeBufSize = zeroBufSize*5 + 30
const benchBufSize = zeroBufSize*100 + 1

func TestMemSetZero(t *testing.T) {
	var data = make([]byte, arraySize)
	_, _ = rand.Read(data)
	SetZero(data)
	for i := range data {
		if data[i] != 0 {
			t.Fatalf("data[%d] should be zero.", i)
		}
	}
}

func TestSetZeroLarge(t *testing.T) {
	var data = make([]byte, largeBufSize)
	_, _ = rand.Read(data)
	SetZero(data)
	for i := range data {
		if data[i] != 0 {
			t.Fatalf("data[%d] should be zero.", i)
		}
	}
}

func TestSetZeroBoundary(t *testing.T) {
	var data = make([]byte, boundaryBufSize)
	_, _ = rand.Read(data)
	SetZero(data)
	for i := range data {
		if data[i] != 0 {
			t.Fatalf("data[%d] should be zero.", i)
		}
	}
}

func setZeroSlicesImpl1(ptr []byte) {
	i := 0
	size := len(ptr)
	for i < size {
		nextSlice := i + zeroBufSize
		if nextSlice > size {
			nextSlice = size
		}
		copy(ptr[i:nextSlice], zeros)
		i = nextSlice
	}
}

func setZeroSlicesImpl2(ptr []byte) {
	i := 0
	size := len(ptr)
	for i < size {
		next := i + zeroBufSize
		copy(ptr[i:], zeros)
		i = next
	}
}

func setZeroSlicesImpl3(ptr []byte) {
	var next int
	size := len(ptr)
	for i := 0; i < size; i = next {
		next = i + zeroBufSize
		copy(ptr[i:], zeros)
	}
}

func BenchmarkSetZeroSlicesImpl1(b *testing.B) {
	var data = make([]byte, benchBufSize)
	for i := 0; i < b.N; i++ {
		setZeroSlicesImpl1(data)
	}
}

func BenchmarkSetZeroSlicesImpl2(b *testing.B) {
	var data = make([]byte, benchBufSize)
	for i := 0; i < b.N; i++ {
		setZeroSlicesImpl2(data)
	}
}

func BenchmarkSetZeroSlicesImpl3(b *testing.B) {
	var data = make([]byte, benchBufSize)
	for i := 0; i < b.N; i++ {
		setZeroSlicesImpl3(data)
	}
}
