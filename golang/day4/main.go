package main

import "fmt"

func main () {
	age := 24

	if age >= 18{
		fmt.Println("Я продам тебе пиво")
	} else {
		if age < 12{
			fmt.Println("Ты чо малой вообще берега попутал пиво покупать?!")
				}else {
					fmt.Println("Алкоголь только с 18-ти лет!")
				}
	}

}


