package main

import (
	"math"
	"runtime/debug"

	"github.com/intervinn/weak"
)

func init() {
	debug.SetGCPercent(-1)
	debug.SetMemoryLimit(math.MaxInt64)
}

func main() {
	vec, err := weak.NewVec[int32](10)
	if err != nil {
		panic(err)
	}

	println(vec.Size())
	vec.Set(0, 69)
	println(vec.At(0))

	vec.Free()
}
