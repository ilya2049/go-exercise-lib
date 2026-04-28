package main

import (
	"fmt"
	"unsafe"
)

func main() {
	slc := []string{}

	for range 2 {
		slc = append(slc, "hello")
	}

	fmt.Print(unsafe.StringData(slc[0]) == unsafe.StringData(slc[1]))
}
