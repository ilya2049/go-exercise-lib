package main

import "fmt"

func main() {
	s := "sun"

	for i, c := range s {
		if i == 1 {
			fmt.Print('d')
		} else {
			fmt.Printf("%c", c)
		}
	}
}
