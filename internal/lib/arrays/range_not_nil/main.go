package main

import "fmt"

func main() {
	var arr [2]int

	for i, e := range arr {
		fmt.Print(i)
		fmt.Print(e)
	}
}
