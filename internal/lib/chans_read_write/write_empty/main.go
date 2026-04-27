package main

import "fmt"

func main() {
	ch := make(chan struct{})

	ch <- struct{}{}

	fmt.Print("OK")
}
