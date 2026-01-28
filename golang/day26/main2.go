package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

type User struct {
	Name    string
	Rating  float64
	Premium bool
}

func main() {

	arrUser := [3]User{
		User{
			Name:    "Даниил",
			Rating:  6.9,
			Premium: true,
		},
		User{
			Name:    "Марина",
			Rating:  8.9,
			Premium: true,
		},
		User{
			Name:    "Алина",
			Rating:  3.5,
			Premium: false,
		},
	}

	fmt.Println("до")
	fmt.Println("--------")
	for i := 0; i < len(arrUser); i++ {
		pp.Println(arrUser[i])
	}
	fmt.Println("")

	for i := 0; i < len(arrUser); i++ {
		if arrUser[i].Premium {
			arrUser[i].Rating += 1.0
		}
	}

	fmt.Println("после")
	fmt.Println("--------")
	for i := 0; i < len(arrUser); i++ {
		pp.Println(arrUser[i])
	}
	fmt.Println("")
}
