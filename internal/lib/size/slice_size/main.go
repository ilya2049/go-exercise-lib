package main

import (
	"fmt"
	"unsafe"
)

func main() {
	bytes := unsafe.Sizeof([]int{1, 2, 3, 4, 5, 6})

	fmt.Print(bytes)
}
