package main 

import (
	"fmt"
)

func main () {
	welcome("саша", 12)
	welcome("петя",13)
	welcome("костя",14)
	welcome("Дарина",15)
}

func welcome (name string, namber int) {
	fmt.Println("добро пожаловать",name,"!")
	fmt.Println("ваш номер",namber, "готов")
	fmt.Println("хорошего отдыха!")
	fmt.Println("")
}

