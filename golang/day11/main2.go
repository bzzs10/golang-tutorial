package main

import(
	"fmt"
)

func main() {

	var chislo int = 10
	var text string = "hello"
	var bools bool = false
	var drob float64 = 2.32

	ukazatel_chislo := &chislo
	ukazatel_text := &text
	ukazatel_bools := &bools
	ukazatel_drob := &drob






	fmt.Println("до вывода",chislo,text,bools,drob)
	fmt.Println("после вывода",ukazatel_chislo,ukazatel_text,ukazatel_bools,ukazatel_drob)
	fmt.Println("после разорхивации",*ukazatel_chislo,*ukazatel_text,*ukazatel_bools,*ukazatel_drob)



}