package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i <= 1_000; i++ {
		time.Sleep(300 * time.Millisecond)
		fmt.Println(i, "-", "german lox:)")
	}
}
