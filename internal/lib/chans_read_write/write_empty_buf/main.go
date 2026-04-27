package main

import "fmt"

func main() {
	ch := make(chan struct{}, 1)

	ch <- struct{}{}

	fmt.Print("OK")
}
