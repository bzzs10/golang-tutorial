package main

import "fmt"

func main() {

	arr := [5]int{5, 15, 61, 54, 1}

	if arr[0]%2 == 0 {
		arr[0] *= 2
	}

	if arr[3]%2 == 0 {
		arr[3] *= 2
	}

	fmt.Println(arr[0])
	fmt.Println(arr[1])
	fmt.Println(arr[2])
	fmt.Println(arr[3])
	fmt.Println(arr[4])
}
