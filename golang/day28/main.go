package main

import (
	"fmt"
	"github.com/k0kubun/pp" 
)

type User struct {	// структурка 
	Name    string
	Rating  float64
	Premium bool
}

func main() {
	UserArray := []User{	//слайс
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
	fmt.Println("len и cap до")	// вывод
	fmt.Println("len",len(UserArray))	// занято на данный момент
	fmt.Println("cap",cap(UserArray))	// кол-во места на данный момент

	UserArray = append(	//append - команда для добавления нового элемента в слайс
		UserArray,
		User{
		Name: "Герман",
		Rating: 8.8,
		Premium: true,
		}, //запятая
	)

	UserArray = append(
		UserArray,
		User{
		Name: "Ваня",
		Rating: 5.6,
		Premium: false,
		},
	)

	fmt.Println("len и cap после")
	fmt.Println("len",len(UserArray))
	fmt.Println("cap",cap(UserArray))

	fmt.Println("до")
	fmt.Println("--------")
	for i := 0; i < len(UserArray); i++ {	// первый вывод элементов из слайса
		pp.Println(UserArray[i])
	}
	fmt.Println("")	// просто пустой отступ

	for i := 0; i < len(UserArray); i++ {	// добавление рейтинга
		if UserArray[i].Premium {
			UserArray[i].Rating += 1.0
		}
	}

	fmt.Println("после")
	fmt.Println("--------")
	for i := 0; i < len(UserArray); i++ {	// второй(итоговый) вывод элементов из слайса
		pp.Println(UserArray[i])
	}
	fmt.Println("") 	// просто пустой отступ
}
