package main

import "fmt"

func main() {
	var x int

	defer func() { fmt.Print(x) }()

	defer func() { x = 1 }()

	defer fmt.Print(x)

	x = 2
}
