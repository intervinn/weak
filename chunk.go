package weak

import (
	"syscall"
)

type Chunk struct {
	Bytes  []byte
	Length int
}

func (c Chunk) Read(buf []byte) (int, error) {
	n := 0
	for i := 0; i < c.Length; i++ {
		buf[i] = c.Bytes[i]
		n++
	}
	return n, nil
}

func (c Chunk) Write(buf []byte) (int, error) {
	n := 0

	for i := 0; i < len(buf); i++ {
		c.Bytes[i] = buf[i]
		n++
	}
	return n, nil
}

func (c Chunk) Free() {
	syscall.Munmap(c.Bytes)
}

func NewChunk(size int) (Chunk, error) {
	fd, err := syscall.Mmap(-1, 0, size,
		syscall.PROT_WRITE|syscall.PROT_READ,
		syscall.MAP_ANON|syscall.MAP_PRIVATE,
	)

	c := Chunk{
		Bytes:  fd,
		Length: len(fd),
	}

	return c, err
}
