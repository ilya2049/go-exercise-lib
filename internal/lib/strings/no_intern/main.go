package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var s1 = string([]byte("hello"))
	const s2 = "hello"

	fmt.Print(unsafe.StringData(s1) == unsafe.StringData(s2))
}
