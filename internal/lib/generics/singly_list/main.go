package main

import "fmt"

func main() {
	for n := fromSlice([]int{1, 2, 3}); n != nil; n = n.Next {
		fmt.Print(n.Value)
	}
}

type Node[T any] struct {
	Next  *Node[T]
	Value T
}

func fromSlice[T any](s []T) *Node[T] {
	var node *Node[T]

	for i := len(s) - 1; i >= 0; i-- {
		node = &Node[T]{Value: s[i], Next: node}
	}

	return node
}
