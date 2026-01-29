package main

import "fmt"

func main() {

	var number int = 10
	var text string = "hello"
	var bools bool = false
	var drob float64 = 2.32

	ptr_number := &number
	ptr_text := &text
	ptr_bools := &bools
	ptr_drob := &drob

	fmt.Println("number до:", number)
	fooint(ptr_number)
	fmt.Println(number)

	fmt.Println("---------------------")

	fmt.Println("text до:", text)
	foostring(ptr_text)
	fmt.Println(text)

	fmt.Println("---------------------")

	fmt.Println("bool до:", bools)
	foobool(ptr_bools)
	fmt.Println(bools)

	fmt.Println("---------------------")

	fmt.Println("drob до:", drob)
	foofloat64(ptr_drob)
	fmt.Println(drob)

}

func foo(ptr *int) {
	*ptr = 15
}

func foo(ptr *string) {
	*ptr = "hi"
}
