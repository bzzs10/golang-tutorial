package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(i, "hello world")

		if i%2 == 0 {
			fmt.Println("четное")
		} else {
			fmt.Println("не четное")
		}
	}
}
