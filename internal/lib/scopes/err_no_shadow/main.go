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
		n, err := getNoErr()

		_, _ = n, err
	}

	return err
}

func getNoErr() (int, error) {
	return 0, nil
}
