package main

import (
	"errors"
	"fmt"
)

func main() {
	err := foo()

	fmt.Print(err == nil)
}

func foo() (err error) {
	bar(func() {
		err = errors.New("e")
	})

	return
}

func bar(fn func()) {
	fn()
}
