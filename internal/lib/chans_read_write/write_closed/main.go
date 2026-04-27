package main

import "fmt"

func main() {
	ch := make(chan struct{}, 1)

	close(ch)

	ch <- struct{}{}

	fmt.Print("OK")
}
