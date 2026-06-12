package main

import "fmt"

func main() {
	array := [...]int{1, 2, 3}

	for i := range array {
		array[i] += 1
	}

	for i := range array {
		fmt.Print(array[i])
	}
}
