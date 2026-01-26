package main

import "fmt"

type car struct {
	HorsePower  int
	marka       string
	model       string
	class       string
	RazgonDoSta float64
}

func newcar(
	HorsePower int,
	marka string,
	model string,
	class string,
	RazgonDoSta float64,
) car {
	if HorsePower < 50 || HorsePower > 100 {
		fmt.Println("horse power не прошла валидацию ")
		return car{}
	}

	if marka == "" {
		fmt.Println("marka не прошла валидацию")
		return car{}
	}

	if model == "" {
		fmt.Println("model не прошла валидацию")
		return car{}
	}

	if class == "" {
		fmt.Println("class не прошел валидацию")
		return car{}
	}

	if RazgonDoSta < 1.0 || RazgonDoSta > 20 {
		fmt.Println("разгон до ста не прошел валидацию")
	}

	return car{
		HorsePower:  HorsePower,
		marka:       marka,
		model:       model,
		class:       class,
		RazgonDoSta: RazgonDoSta,
	}
}

func (c Car) HorsePowerUP(HorsePower int) {

}

func main() {
	// car1 := car{
	// 	HorsePower:  40,
	// 	marka:       "haval",
	// 	model:       "jolion",
	// 	class:       "Кросовер",
	// 	RazgonDoSta: 13.4,
	// }

	// car2 := car{
	// 	HorsePower:  51,
	// 	marka:       "BMW",
	// 	model:       "x5",
	// 	class:       "Седан",
	// 	RazgonDoSta: 5.6,
	// }

	// fmt.Println(car1)
	// fmt.Println(car2)

	car := newcar(100, "honda", "civic", "седан", 20.0)
	fmt.Println(car)
}
