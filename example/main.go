package main

import (
	"runtime/debug"

	"github.com/intervinn/weak"
)

func init() {
	debug.SetGCPercent(-1)
}

type Container struct {
	Value int32
}

func main() {
	buf, err := weak.Alloc[Container](weak.SizeOf(Container{}))
	if err != nil {
		panic(err)
	}
	defer buf.Free()
	buf.Write(Container{
		Value: 69420,
	})

	c, err := buf.Read()
	if err != nil {
		panic(err)
	}
	println(c.Value)
}
