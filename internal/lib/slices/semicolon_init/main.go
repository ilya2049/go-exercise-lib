package main

import "fmt"

func main() {
	var numbers = []int{
		1: 2,
		3: 4,
	}

	for _, n := range numbers {
		fmt.Print(n)
	}
}
