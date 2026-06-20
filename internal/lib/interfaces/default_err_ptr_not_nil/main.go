package main

import "fmt"

type Error struct {
}

func (e Error) Error() string {
	return "error"
}

func main() {
	err := foo()

	fmt.Print(err == nil)
}

func foo() error {
	var p *Error = nil

	return p
}
