package main

import (
	"fmt"
	"strconv"
)

type Error struct {
	code int
}

func (e Error) Error() string {
	return "invalid code: " + strconv.Itoa(e.code)
}

func main() {
	a := Error{code: 1}
	b := Error{code: 1}

	fmt.Print(error(a) == error(b))
}
