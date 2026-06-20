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
	defer func() { err = catch(recover()) }()

	panic("")
}

func catch(r any) error {
	if r != nil {
		return errors.New("")
	}

	return nil
}
