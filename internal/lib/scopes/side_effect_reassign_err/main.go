package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("main")

	x, err := foo()
	_ = x

	fmt.Print(err)
}

func foo() (int, error) {
	return 1, errors.New("foo")
}
