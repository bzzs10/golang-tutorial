package main

import (
	"fmt"
	"math/rand"
	"time"
)

	var score int = 0

func main() {


	fmt.Println("--------------------------")
	fmt.Println("начать игру")
	fmt.Println("ваш счет:", score)

	for text := 1; text <= 5; text++ {

		time.Sleep(100 * time.Millisecond)

		fmt.Println("--------------------------")
		fmt.Println("вы совершили прыжок, вам дали +5 очков")
		score = score + 5
		fmt.Println("ваш счет:", score)

		if rand.Intn(2) == 1 {
			fmt.Println("--------------------------")
			fmt.Println("упс.. вы врезались")
			fmt.Println("ваш счет:", score)
			break
		}
		fmt.Println("--------------------------")
		fmt.Println("вы совершили прыжок, вам дали +5 очков")
		fmt.Println("ваш счет:", score)

		// return score
	}	

	end()
}

func end() {
	fmt.Println("--------------------------")
	fmt.Println("Game over")
	fmt.Println("ваш счет:",score)
}