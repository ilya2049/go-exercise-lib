package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a [17]byte

	fmt.Print(unsafe.Sizeof(a))
}
