package main

import "fmt"

func main() {
	var slc = []int{1, 2, 3}

	for i, e := range slc {
		if i == 0 {
			slc[1] = 11
		}

		fmt.Print(e)
	}
}
