package main

import (
	"errors"
	"fmt"
)

func main() {
	err := foo()

	fmt.Print(err == nil)
}

func bar() error {
	return errors.New("foo")
}

func foo() (err error) {
	if err := bar(); err == nil {
	}

	return
}
