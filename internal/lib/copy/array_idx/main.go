package main

import "fmt"

const N = 3

func main() {
	array := [N]int{1, 2, 3}

	for i := range array {
		array[i] += 1
	}

	for i := range array {
		fmt.Print(array[i])
	}
}
