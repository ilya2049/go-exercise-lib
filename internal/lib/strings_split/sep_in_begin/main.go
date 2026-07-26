package main

import (
	"fmt"
	"strings"
)

func main() {
	result := ""

	for s := range strings.SplitSeq("/a/b/c", "/") {
		if s == "" {
			result += "?"
		} else {
			result += s
		}
	}

	fmt.Print(result)
}
