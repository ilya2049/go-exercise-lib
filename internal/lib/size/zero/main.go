package main

import (
	"fmt"
	"unsafe"
)

type A [0][10]int

type B struct {
	x A
	y [3]A
	z [4]struct{}
}

func main() {
	fmt.Print(unsafe.Sizeof(B{}))
}
