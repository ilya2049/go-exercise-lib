package main

import (
	"fmt"
	"unsafe"
)

type X struct {
	a int
	b int
	c int
}

func (x *X) Foo() {
	fmt.Print(unsafe.Sizeof(x))
}

func main() {
	x := X{}
	x.Foo()
}
