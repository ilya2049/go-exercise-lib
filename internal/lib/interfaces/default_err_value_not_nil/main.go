package main

import "fmt"

type Error struct {
}

func (e Error) Error() string {
	return "error"
}

func main() {
	var err Error

	foo(err)
}

func foo(err error) {
	fmt.Print(err == nil)
}
