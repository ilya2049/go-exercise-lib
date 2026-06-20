package main

import "fmt"

const (
	a float64 = iota
	b
	c
)

func main() {
	fmt.Print(a)
	fmt.Print(b)
	fmt.Print(c)
}
