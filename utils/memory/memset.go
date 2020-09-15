package memory

// #include <string.h>
import "C"
import "unsafe"

const zeroBufSize = 4 * 1024

var zeros = make([]byte, zeroBufSize)

func setZeroC(ptr []byte) {
	C.memset(unsafe.Pointer(&ptr[0]), C.int(0), C.ulonglong(len(ptr)))
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
