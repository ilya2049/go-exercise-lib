package main

import "fmt"

func main() {
	var (
		v   int   = 100
		pt1 *int  = &v
		pt2 **int = &pt1
	)

	*pt1 = 200
	fmt.Print(**pt2)
}
