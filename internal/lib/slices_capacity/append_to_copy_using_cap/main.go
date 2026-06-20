package main

import "fmt"

func main() {
	numbers := make([]int, 0, 5)
	numbers = append(numbers, 1, 2, 3)

	add(numbers, 4)

	for _, n := range numbers[:4] {
		fmt.Print(n)
	}
}

func add(numbers []int, n int) {
	numbers = append(numbers, n)
}
