package main

import "fmt"

func main() {
	x := 1

	fmt.Print(x)

	{
		fmt.Print(x)

		x := 2

		fmt.Print(x)
	}

	fmt.Print(x)
}
