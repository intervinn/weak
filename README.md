implements manual memory allocation using syscalls, memory leaks guaranteed!

almost everything in the library is stack-allocated except for binary endings used to turn values into byte arrays,
but some day ill find a better alternative


## usage

### chunks 
chunks are just a wrapping around allocated memory which you can write and read from

unlike other implementations of io.Reader/io.Writers, it doesn't append to slice, 
but writes from beginning
```go
package main

import (
	"encoding/binary"

	"github.com/intervinn/weak"
)

func main() {
	c, err := weak.Alloc(weak.SizeOf[int32]())
	if err != nil {
		panic(err)
	}

	var val int32
	binary.Write(c, binary.NativeEndian, int32(69420)) // 69420
	binary.Read(c, binary.NativeEndian, &val)
	println(val)

	c.Free()
	//binary.Read(c, binary.NativeEndian, &val) this segfaults
}
```

### boxes
boxes wrap around chunks to hold a single instance of datatype
```go
import "github.com/intervinn/weak"

func main() {
	b, err := weak.NewBox[float32]()
	if err != nil {
		panic(err)
	}

	b.Set(3.14)
	println(b.Value()) // +3.140000e+000 but thats println issue

	b.Free()
}
```