package main

import "fmt"

func main() {
	var ch chan bool = nil

	for range ch {
		fmt.Print("OK")
	}
}
