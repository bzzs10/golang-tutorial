package main

import (
	"fmt"
)

func main() {
	text1 := san(1, 4)
	fmt.Println(text1)
	fmt.Println(bob())
}

func bob() string {
	return "hello"
}

func san(a int, b int) int {
	s := a + b
	return s
}
	