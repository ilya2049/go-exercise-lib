package main

import "fmt"

func main() {
	for v := range iter {
		fmt.Print(v)

		break
	}
}

func iter(yield func(string) bool) {
	yield("a")
	yield("b")
	yield("c")
}
