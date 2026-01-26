package main

import(
	"fmt"
)

var drob float64 = 10.66

func main () {
	fmt.Println("1", drob)
	z()

	fmt.Println("2",drob)
	x()

	fmt.Println("3",drob)
	c()

}

func z() {
	drob *= 2 
}

func x() {
	 drob *= 19
}


func c() {
	drob *= 12
}