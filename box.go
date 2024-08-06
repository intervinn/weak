package weak

import "encoding/binary"

type Box[T any] struct {
	Chunk Chunk
}

func (b Box[T]) Compare(other Box[T]) bool {
	if other.Chunk.Length != b.Chunk.Length {
		return false
	}

	for i := 0; i < b.Chunk.Length; i++ {
		if b.Chunk.Bytes[i] != other.Chunk.Bytes[i] {
			return false
		}
	}
	return true
}

func (b Box[T]) Value() T {
	var res T
	binary.Read(b.Chunk, binary.NativeEndian, &res)
	return res
}

func (b Box[T]) Set(v T) {
	binary.Write(b.Chunk, binary.NativeEndian, v)
}

func NewBox[T any]() (Box[T], error) {
	c, err := NewChunk(SizeOf[T]())

	return Box[T]{
		Chunk: c,
	}, err
}

func (b Box[T]) Free() {
	b.Chunk.Free()
}
