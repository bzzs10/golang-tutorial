package main

import "fmt"

func main() {
	intSlice := make([]int, 0, 5)

	intSlice = append(intSlice, 10, 15, 25, 150)
	fmt.Println(len(intSlice))
	fmt.Println(cap(intSlice))
	fmt.Println(intSlice)

}
