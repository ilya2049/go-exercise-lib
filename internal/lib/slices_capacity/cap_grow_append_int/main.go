package main

import "fmt"

func main() {
	a := []int{1, 2, 3}

	c := append(a, 4)

	fmt.Print(cap(c))
}
