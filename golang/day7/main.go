package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	score := 0

	for {

		fmt.Println("-------------------------")
		fmt.Println("я в пролетел через трубу")
		score = score + 1
		fmt.Println("ваш счет", score)

		time.Sleep(100 * time.Millisecond)

		fmt.Println("-------------------------")
		if rand.Intn(9) == 1 {
			fmt.Println("я резался в трубу")
			fmt.Println("game over!")
			fmt.Println("ваш счет", score)
			fmt.Println("-------------------------")
			break
		}

	}
}
