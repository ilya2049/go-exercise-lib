package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}
	s := numbers[:2]

	for _, n := range s {
		fmt.Print(n)
	}
}
