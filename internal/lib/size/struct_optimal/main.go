package main

import (
	"fmt"
	"unsafe"
)

type A struct {
	a int64
	b bool
	c bool
}

func main() {
	fmt.Print(unsafe.Sizeof(A{}))
}
