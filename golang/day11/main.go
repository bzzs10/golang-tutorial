package main

import(
	"fmt"
)

func main() {
	name := "Игорь"

	ptr :=&name

	fmt.Println("имя до", name)

	chName(ptr)

	fmt.Println("имя после", name)




}


func chName(s *string) {
	*s = "иван"
}