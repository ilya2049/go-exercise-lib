package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Print(foo() == nil)
}

func foo() error {
	err := errors.New("error")
	if err != nil {
		err = getNoErr()

		_ = err
	}

	return err
}

func getNoErr() error {
	return nil
}
