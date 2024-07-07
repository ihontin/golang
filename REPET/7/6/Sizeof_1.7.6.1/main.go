package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var (
		bo         bool
		sizeInt    int
		sizeInt8   int8
		sizeInt16  int16
		sizeInt32  int32
		sizeInt64  int64
		sizeUint   uint
		sizeUint8  uint8
		sizeStruct = struct{}{}
	)

	fmt.Printf("bool size: %d\n", int(unsafe.Sizeof(bo)))
	fmt.Printf("int size: %d\n", int(unsafe.Sizeof(sizeInt)))
	fmt.Printf("int8 size: %d\n", int(unsafe.Sizeof(sizeInt8)))
	fmt.Printf("int16 size: %d\n", int(unsafe.Sizeof(sizeInt16)))
	fmt.Printf("int32 size: %d\n", int(unsafe.Sizeof(sizeInt32)))
	fmt.Printf("int64 size: %d\n", int(unsafe.Sizeof(sizeInt64)))
	fmt.Printf("uint size: %d\n", int(unsafe.Sizeof(sizeUint)))
	fmt.Printf("uint8 size: %d\n", int(unsafe.Sizeof(sizeUint8)))
	fmt.Printf("struct size: %d\n", int(unsafe.Sizeof(sizeStruct)))

}
