package main

import "fmt"

type X struct {
	s []float64
}

func NewX() X {
	return X{
		s: make([]float64, 0, 1),
	}
}

func (x *X) append(n float64) {
	x.s = append(x.s, n)
}

func main() {
	x := NewX()

	x.append(1)

	fmt.Print(x.s[0])
}
