package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s := "Hello, world!"

	fmt.Print(unsafe.Sizeof(s))
}
