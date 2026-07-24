package main

import "fmt"

func main() {
	var m map[int]map[string]bool

	k := m[1]["x"]

	fmt.Print(k)
}
