package main

import "fmt"

func main() {
	var src, dst []int

	src = []int{1, 2, 3}
	dst = make([]int, 0, len(src))

	copy(dst, src)

	for _, n := range dst {
		fmt.Print(n)
	}
}
