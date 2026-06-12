package main

import "fmt"

type User struct {
	Age int
}

func main() {
	users := []*User{
		{Age: 6},
		{Age: 7},
	}

	for _, u := range users {
		u.Age += 1
	}

	for _, u := range users {
		fmt.Print(u.Age)
	}
}
