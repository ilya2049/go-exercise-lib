package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var s1 = "hello"
	const s2 = "hello"

	fmt.Print(unsafe.StringData(s1) == unsafe.StringData(s2))
}
