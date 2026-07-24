package main

import "fmt"

func main() {
	x := 1

	switch x {
	case 0:
		fmt.Print("0")
	case 1:
	case 2:
	default:
		fmt.Print("3")
	}
}
