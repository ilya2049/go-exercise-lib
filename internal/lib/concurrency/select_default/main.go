package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)

	ch := make(chan int)

	go func() {
		for i := 1; i <= 3; i++ {
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
			runtime.Gosched()
		}
	}
}
