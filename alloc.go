package weak

import (
	"bytes"
	"encoding/binary"
	"syscall"
)

type Buffer[T any] struct {
	buf *bytes.Buffer
	len int
}

// Allocates memory and creates a new buffer
func Alloc[T any](len int) (*Buffer[T], error) {
	fd, err := syscall.Mmap(-1, 0, len,
		syscall.PROT_WRITE|syscall.PROT_READ,
		syscall.MAP_ANONYMOUS|syscall.MAP_PRIVATE,
	)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(fd[:0])
	return &Buffer[T]{
		buf: buf,
		len: len,
	}, nil
}

// Read the value written into the buffer
func (b *Buffer[T]) Read() (T, error) {
	var res T
	err := binary.Read(b.buf, binary.NativeEndian, &res)
	return res, err
}

// Write value into buffer
func (b *Buffer[T]) Write(v T) error {
	return binary.Write(b.buf, binary.NativeEndian, v)
}

// Free the allocated memory
func (b *Buffer[T]) Free() {
	syscall.Munmap(b.buf.Bytes())
}
