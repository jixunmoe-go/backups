package memory

// #include <string.h>
// void ClearMem(void* ptr, int size) {memset(ptr, 0, (size_t)size);}
import "C"
import "unsafe"

const zeroBufSize = 4 * 1024

var zeros = make([]byte, zeroBufSize)

func setZeroC(ptr []byte) {
	C.ClearMem(unsafe.Pointer(&ptr[0]), C.int(len(ptr)))
}

// SetZero is the fastest implementation to zero out a given slice.
func SetZero(ptr []byte) {
	if len(ptr) > zeroBufSize {
		// Slow
		setZeroC(ptr)
	} else {
		copy(ptr, zeros)
	}
}
