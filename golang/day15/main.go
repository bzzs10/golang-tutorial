package main

import (
	"fmt"
)

type User struct {
	Name       string
	Age        int
	PhoneNuber string
	IsClose    bool
	Rating     float64
}

func main() {
	user := User{
		Name:       "bob",
		Age:        25,
		PhoneNuber: "+9(545)210-23-03",
		IsClose:    true,
		Rating:     3.4,
	}

	hello(user)
	fmt.Println()

	fmt.Println("full-user:", user)

	fmt.Println("name:", user.Name)
	fmt.Println("Age:", user.Age)
	fmt.Println("PhoneNuber:", user.PhoneNuber)
	fmt.Println("IsClose:", user.IsClose)
	fmt.Println("IsClose:", user.Rating)

	fmt.Println("--------------------")

	fmt.Println("full-user:", user)

	user.Name = "kitty"
	fmt.Println("name:", user.Name)
	user.Age = 900
	fmt.Println("Age:", user.Age)
	user.PhoneNuber = "+1(909)09-09-999"
	fmt.Println("PhoneNuber:", user.PhoneNuber)
	user.IsClose = false
	fmt.Println("IsClose:", user.IsClose)
	user.Rating = 99.9
	fmt.Println("IsClose:", user.Rating)

}

func hello(u User) {
	fmt.Println("всем привет меня зовут!", u.Name)
	fmt.Println("мой рейтинг:", u.Rating)
}
