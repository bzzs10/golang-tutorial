package main

import "fmt"


var nacalo string = "начало"
var end string = "конец"


func main() {
    amongus := "defer1"


    defer func() {
        fmt.Println(amongus)
    }()

    fmt.Println("hello1")
    fmt.Println("hello2")
    fmt.Println("hello3")

	marina()
	german()
}

func marina() {

	fmt.Println(nacalo)
	defer func() {
		fmt.Println(end)
	}()

	fmt.Println("некий текст, в функции названой в честь Марины"sefsefp[SE{OPFK<os[keofk[ofoksa[kef f]]]}])
	fmt.Println("данная фукнция идет как эксперемент для закрпеления темы defer")
}

func german() {

	fmt.Println(nacalo)
	defer func() {
		fmt.Println(end)
	}()

	fmt.Println("некий текст, в функции названой в честь German")
	fmt.Println("данная фукнция идет как эксперемент для закрпеления темы defer")
}
