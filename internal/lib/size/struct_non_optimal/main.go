package main

import (
	"fmt"
	"unsafe"
)

type A struct {
	a bool
	b int64
	c bool
}

func main() {
	fmt.Print(unsafe.Sizeof(A{}))
}
