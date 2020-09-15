package memory

import "testing"

var arraySize = 32

func BenchmarkMemSet(b *testing.B) {
	var data = make([]byte, arraySize)
	for i := 0; i < b.N; i++ {
		setZeroC(data)
	}
}

func BenchmarkMemSetZero(b *testing.B) {
	var data = make([]byte, arraySize)
	for i := 0; i < b.N; i++ {
		SetZero(data)
	}
}
