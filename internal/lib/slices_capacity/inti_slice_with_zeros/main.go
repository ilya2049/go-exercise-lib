package main

import "fmt"

func main() {
	numbers := make([]int, 1)

	numbers = append(numbers, 2)

	for _, n := range numbers {
		fmt.Print(n)
	}
}
