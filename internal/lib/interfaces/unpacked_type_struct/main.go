package main

import "fmt"

type Error struct {
}

func (e Error) Error() string {
	return "error"
}

func main() {
	err := Error{}

	foo(err)
}

func foo(err error) {
	switch err.(type) {
	case Error:
		fmt.Print("struct")
	case *Error:
		fmt.Print("pointer")
	default:
		fmt.Print("default")
	}
}
