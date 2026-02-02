package main

import "fmt"

type Car struct {
	marka  string
	speed  int
	engine float64
}

type myslise struct {
	Number int
	String string
}

func main() {
	masiw_int := [5]int{1, 2, 3, 4, 5}
	masiw_sting := [3]string{"Саша", "Маша", "Петя"}
	masiw_float := [5]float64{1.6, 15.7, 26.3}
	masiw_bool := [1]bool{} // можно ничего не класть по сути...(значение по умолчанию меня устраивает)

	fmt.Println("-------масивы--------")
	fmt.Println(masiw_int)
	fmt.Println(masiw_sting)
	fmt.Println(masiw_float)
	fmt.Println(masiw_bool)
	fmt.Println("--------вывод отдельного индекса--------")
	fmt.Println(masiw_sting[1])
	fmt.Println(masiw_int[1])
	fmt.Println("--------слайсы--------")

	structur := make([]myslise, 0, 10)

	Carbaze := []Car{
		Car{
			marka:  "lADA",
			speed:  100,
			engine: 1.7,
		},
		Car{
			marka:  "AUDI",
			speed:  500,
			engine: 4.8,
		},
		Car{
			marka:  "BMW",
			speed:  140,
			engine: 2.76,
		},
	}

	Carbaze = append(
		Carbaze,
		Car{
			marka:  "MERSEDES",
			speed:  250,
			engine: 2.51,
		},
	)

	fmt.Println(Carbaze)
	fmt.Println("мерседес был выведен отдельно")
	fmt.Println("--------maps-----------")

	contact := map[string]int{
		"Марина":      79595129358,
		"Алиса":       88555353554,
		"Дима":        65689150142,
		"Ви":          12341241245,
		"Сесра Алисы": 51531513515,
	}

	fmt.Println(contact)
	fmt.Println("---------------")
	fmt.Println(contact["Марина"])

}
