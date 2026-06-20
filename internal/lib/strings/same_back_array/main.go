package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s := "dark castle"

	b := s[:3]

	p1 := unsafe.StringData(s)
	p2 := unsafe.StringData(b)

	fmt.Print(p1 == p2)
}
