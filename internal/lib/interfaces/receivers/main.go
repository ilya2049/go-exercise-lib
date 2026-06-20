package main

import "fmt"

type InError struct {
}

func (e InError) Error() string {
	return "error"
}

type OutError struct {
}

func (e *OutError) Error() string {
	return "error"
}

func main() {
	b := equal(InError{}, &OutError{})

	// What lines will not compile?
	// b := equal(&InError{}, &OutError{})
	// b := equal(&InError{}, OutError{})
	// b := equal(InError{}, OutError{})

	fmt.Print(b)
}

func equal(err1, err2 error) bool {
	return err1.Error() == err2.Error()
}
