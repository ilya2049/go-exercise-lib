package main

import "fmt"

func main() {
	numbers := [2]int{1, 2}
	add(numbers)

	for _, n := range numbers {
		fmt.Print(n)
	}
}

func add(s [2]int) {
	s[0] = 3
}
