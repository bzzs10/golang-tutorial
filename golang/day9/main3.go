package main 

import (
	"fmt"
	"math"
)

func main () {

	a := 10
	b := 123
	c := 5

	square(a, b, c)
}


func square(a int, b int, c int) {

	d := b*b - 4*a*c

 	if d < 0 {
		fmt.Println("корней нету")
		return
	}

	if d == 0 {
		x := -b / (2 * a)

		fmt.Println("у уровнения 1 корень", x)

		return
	}

	x1 := (float64(-b) + math.Sqrt(float64(d))) / float64(2*a)
	x2 := (float64(-b) - math.Sqrt(float64(d))) / float64(2*a)
	
	fmt.Println("первый корень", x1)
	fmt.Println("второй корень", x2)
}
