package main

import "fmt"

const N = 3

func main() {
	array := [N]int{1, 2, 3}

	for _, e := range array {
		e += 1
	}

	for _, e := range array {
		fmt.Print(e)
	}
}
