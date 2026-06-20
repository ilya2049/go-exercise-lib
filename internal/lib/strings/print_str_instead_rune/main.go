package main

import "fmt"

func main() {
	s := "sun"

	for i, c := range s {
		if i == 2 {
			fmt.Print("m")
		} else {
			fmt.Printf("%c", c)
		}
	}
}
