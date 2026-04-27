package detect

import "fmt"

func Panic(fn func()) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Print("panic")
		}
	}()
	fn()
}
