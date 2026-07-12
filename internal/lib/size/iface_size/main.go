package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i any = nil

	fmt.Print(unsafe.Sizeof(i))
}
