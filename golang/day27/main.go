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
	UserArray := []User{
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
	fmt.Println("len и cap до")
	fmt.Println("len",len(UserArray))
	fmt.Println("cap",cap(UserArray))

	UserArray = append(
		UserArray,
		User{
		Name: "Герман"
		Rating: 8.8,
		Premium: true,
		}
	)
	fmt.Println("len и cap после")
	fmt.Println("len",len(UserArray))
	fmt.Println("cap",cap(UserArray))

	fmt.Println("до")
	fmt.Println("--------")
	for i := 0; i < len(UserArray); i++ {
		pp.Println(UserArray[i])
	}
	fmt.Println("")

	for i := 0; i < len(UserArray); i++ {
		if UserArray[i].Premium {
			UserArray[i].Rating += 1.0
		}
	}

	fmt.Println("после")
	fmt.Println("--------")
	for i := 0; i < len(UserArray); i++ {
		pp.Println(UserArray[i])
	}
	fmt.Println("")
}