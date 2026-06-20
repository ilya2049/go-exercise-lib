package main

import "fmt"

func main() {
	a := 1

	for _, a := range []int{2, 3} {
		fmt.Print(a)
	}

	fmt.Print(a)
}
