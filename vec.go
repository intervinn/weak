package weak

import (
	"bytes"
	"encoding/binary"
)

// A dynamic array implementation
type Vec[T any] struct {
	size  int64
	isize int64
	buf   *bytes.Buffer
}

// Constructs new Vec
func NewVec[T any](size int) (Vec[T], error) {
	b, err := Alloc(SizeOf[T]() * size)

	v := Vec[T]{
		size:  int64(size),
		isize: int64(SizeOf[T]()),
		buf:   b,
	}
	return v, err
}

func (v Vec[T]) Size() int64 {
	return v.size
}

func (v Vec[T]) Set(index int64, value T) {
	bgn := index * v.isize
	tmp := bytes.NewBuffer([]byte{})

	bt := v.buf.Bytes()

	binary.Write(tmp, binary.NativeEndian, value)
	for i, v := range tmp.Bytes() {
		bt[bgn+int64(i)] = v
	}
}

func (v Vec[T]) At(index int64) T {
	bgn := index * v.isize
	bt := v.buf.Bytes()
	var res T

	tmp := bytes.NewBuffer(bt[bgn : bgn+v.isize])
	binary.Read(tmp, binary.NativeEndian, &res)

	return res
}

func (v Vec[T]) Free() {
	Free(v.buf)
}
