package weak

import (
	"bytes"
	"syscall"
)

// Allocates memory for later use with io.Readers and io.Writers
func Alloc(len int) (*bytes.Buffer, error) {
	fd, err := syscall.Mmap(-1, 0, len,
		syscall.PROT_WRITE|syscall.PROT_READ,
		syscall.MAP_ANONYMOUS|syscall.MAP_PRIVATE,
	)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(fd)
	return buf, nil
}

func Free(b *bytes.Buffer) {
	syscall.Munmap(b.Bytes())
}
