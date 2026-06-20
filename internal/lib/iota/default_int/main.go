package main

import (
	"fmt"
	"reflect"
)

const (
	a = iota
	b
)

func main() {
	fmt.Print(reflect.TypeOf(b).String())
}
