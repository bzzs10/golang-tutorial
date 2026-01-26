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

	fmt.Println(arr)
	// fmt.Println(arr[1])
	// fmt.Println(arr[2])
	// fmt.Println(arr[3])
	// fmt.Println(arr[4])

	arr2 := [7]int{}
	// arr2[0] = arr[0]
	// arr2[1] = arr[1]
	// arr2[2] = arr[2]
	// arr2[3] = arr[3]
	// arr2[4] = arr[4]
	// arr2[5] = 7

	fmt.Println(arr2)
}
