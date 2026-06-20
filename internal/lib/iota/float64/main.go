package main

import (
	"fmt"
	"reflect"
)

const (
	a float64 = iota
	b
)

func main() {
	fmt.Print(reflect.TypeOf(b).String())
}
