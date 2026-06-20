package main

import "fmt"

func main() {
	numbers := []int{1, 2}
	add(&numbers)

	for _, n := range numbers {
		fmt.Print(n)
	}
}

func add(n *[]int) {
	*n = append(*n, 3)
}
