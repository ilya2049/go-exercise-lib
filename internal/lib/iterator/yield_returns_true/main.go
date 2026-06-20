package main

import "fmt"

func main() {
	for v := range iterate {
		fmt.Print(v)
	}
}

func iterate(yield func(string) bool) {
	if yield("a") {
		if yield("b") {
			yield("c")
		}
	}
}
