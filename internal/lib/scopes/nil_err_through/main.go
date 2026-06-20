package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("e")

	err = foo(func(err error) error {
		if err != nil {
			return err
		}

		return nil
	})

	fmt.Print(err == nil)
}

func foo(fn func(err error) error) (err error) {
	err = fn(err)

	return
}
