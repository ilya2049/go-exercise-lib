package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}
	b := []int{4, 5}

	c := append(a, b...)

	fmt.Print(cap(c))
}
