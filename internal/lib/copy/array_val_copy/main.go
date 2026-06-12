package main

import "fmt"

func main() {
	array := [...]int{1, 2, 3}

	for _, e := range array {
		e += 1
	}

	for _, e := range array {
		fmt.Print(e)
	}
}
