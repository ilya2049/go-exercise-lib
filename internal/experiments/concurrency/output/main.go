package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := range 3 {
		wg.Go(func() {
			fmt.Print(i)
		})
	}

	wg.Wait()
}
