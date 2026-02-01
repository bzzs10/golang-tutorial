package main

import "fmt"

func main() {
	weather := map[int]int{
		11: +3,
		12: +6,
		13: +9,
		14: -4,
		15: +1,
		// 30: 0,
	}
	c, ok := weather[30]
	fmt.Println(weather[14])
	fmt.Println(c, ok)
	fmt.Println(weather[30])

	for k, _ := range weather {
		weather[k] += 1
	}
	weather[30] = -30
	fmt.Println(weather)
}
