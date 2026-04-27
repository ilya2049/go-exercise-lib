package detect

import (
	"fmt"
	"time"
)

func Deadlock(fn func()) {
	c := make(chan struct{})

	go func() {
		fn()
		c <- struct{}{}
	}()

	select {
	case <-c:
		fmt.Print("OK")
	case <-time.After(50 * time.Millisecond):
		fmt.Print("deadlock")
	}
}
