package main

import "fmt"

func main() {
	numbers := make([]int, 0, 2)

	numbers[0] = 1
	numbers[1] = 2

	for _, n := range numbers[:4] {
		fmt.Print(n)
	}
}
