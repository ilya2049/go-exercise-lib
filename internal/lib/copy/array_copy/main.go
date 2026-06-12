package main

import "fmt"

func main() {
	var array = [...]int{1, 2, 3}

	for i, e := range array {
		if i == 0 {
			array[1] = 11
		}

		fmt.Print(e)
	}
}
