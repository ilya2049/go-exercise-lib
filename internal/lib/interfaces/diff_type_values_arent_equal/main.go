package main

import (
	"fmt"
	"strconv"
)

type ErrorA struct {
	code int
}

func (e ErrorA) Error() string {
	return "invalid code: " + strconv.Itoa(e.code)
}

type ErrorB struct {
	code int
}

func (e ErrorB) Error() string {
	return "invalid code: " + strconv.Itoa(e.code)
}

func main() {
	a := ErrorA{code: 1}
	b := ErrorB{code: 1}

	fmt.Print(error(a) == error(b))
}
