package main

import "fmt"

func main() {
	for v := range iter {
		fmt.Print(v)
	}
}

func iter(yield func(int) bool) {
	_ = yield(1) && yield(2) && yield(3)
}
