package main

import "fmt"

func main() {
	arr := [7]int{5, 4, 1, 6, 88, 8}

	for i := 0; i < 7; i++ {
		if arr[i]%2 == 0 {
			arr[i] *= 2
		}
	}

	for i := 0; i < 7; i++ {
		fmt.Println(i, "-", arr[i])
	}

}
