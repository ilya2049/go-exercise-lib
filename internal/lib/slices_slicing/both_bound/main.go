package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	s := numbers[1:3]

	for _, n := range s {
		fmt.Print(n)
	}
}
