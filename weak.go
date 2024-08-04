package weak

import "unsafe"

// A shorthand for getting type size
func SizeOf[T any]() int {
	var a T
	return int(unsafe.Sizeof(a))
}
