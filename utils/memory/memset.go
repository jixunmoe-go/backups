package memory

const zeroBufSize = 1024

var zeros = make([]byte, zeroBufSize)

// SetZero can clear a given byte slice and reset all values to zero.
func SetZero(ptr []byte) {
	var nextSlice int
	size := len(ptr)
	for i := 0; i < size; i = nextSlice {
		nextSlice = i + zeroBufSize
		copy(ptr[i:], zeros)
	}
}
