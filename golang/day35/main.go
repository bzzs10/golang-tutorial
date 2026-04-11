package main

import (
	"fmt"
)

type Auto interface {
	stepOnGas()
	stepOnBrake()
	bibib()
}

type Zhiga struct{}

func (z Zhiga) stepOnGas() {
	fmt.Println("Я жига постараюсь не развалиться!")
}

func (z Zhiga) stepOnBrake() {
	fmt.Println("страшный свист колодок")
}

func (z Zhiga) bibib() {
	fmt.Println("класкосн не рабоатает)))")
}

type lamba struct{}

func (l lamba) stepOnGas() {
	fmt.Println("я ламба, оч быстрая!")
}

func (l lamba) stepOnBrake() {
	fmt.Println("резкий тормоз! дорогой скрик колодок")
}

func (l lamba) bibib() {
	fmt.Println("Раскошный бибибиб")
}

func ride(auto Auto) {
	fmt.Println("я водитель!")
	fmt.Println("я сажусь в свою машину!")
	fmt.Println("и нажимаю на газ...")
	auto.stepOnGas()
	auto.stepOnBrake()
	auto.bibib()
}

func main() {

	NameAuto := "1"

	if NameAuto == "1" {
		lamba := lamba{}
		ride(lamba)
	} else {
		if NameAuto == "2" {
			Zhiga := Zhiga{}
			ride(Zhiga)
		} else {
			return
		}
	}

}
