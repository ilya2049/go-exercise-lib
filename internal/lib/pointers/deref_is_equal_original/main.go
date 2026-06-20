package main

import "fmt"

func main() {
	type X struct{}

	x := X{}
	y := &x
	z := *y

	fmt.Print(x == z)
}
