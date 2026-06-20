package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	for n := range numbers {
		fmt.Print(n)
	}
}
