package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 2; i++ {
			ch <- i
		}

		close(ch)
	}()

loop:
	for {
		select {
		case value, ok := <-ch:
			if !ok {
				break loop
			}

			fmt.Print(value)
		default:
			fmt.Print("?")
		}
	}
}
