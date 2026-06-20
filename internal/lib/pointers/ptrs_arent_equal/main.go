package main

import "fmt"

func main() {
	type X struct{}

	a := &X{}
	b := *a
	c := &b

	fmt.Print(c == a)
}
