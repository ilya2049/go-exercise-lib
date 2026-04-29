package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}
	s := numbers[1:2:2]

	s[0] = 7

	for _, n := range numbers {
		fmt.Print(n)
	}
}
