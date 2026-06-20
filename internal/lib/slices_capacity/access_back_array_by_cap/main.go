package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	s1 := numbers[:2]
	s2 := s1[:len(s1)+2]

	for _, n := range s2 {
		fmt.Print(n)
	}
}
