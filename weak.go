package weak

import "unsafe"

// A shorthand for int(unsafe.Sizeof(A))
func SizeOf(a any) int {
	return int(unsafe.Sizeof(a))
}
