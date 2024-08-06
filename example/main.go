package main

import (
	"encoding/binary"

	"github.com/intervinn/weak"
)

func main() {
	c, err := weak.NewChunk(weak.SizeOf[int32]())
	if err != nil {
		panic(err)
	}

	var val int32
	binary.Write(c, binary.NativeEndian, int32(69420)) // 69420
	binary.Read(c, binary.NativeEndian, &val)
	println(val)

	c.Free()
	//binary.Read(c, binary.NativeEndian, &val) this segfaults

	b, err := weak.NewBox[float32]()
	if err != nil {
		panic(err)
	}

	b.Set(3.14)
	println(b.Value()) // +3.140000e+000 but thats println issue

	b.Free()
}
