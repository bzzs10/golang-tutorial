package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("Введите комануд: ")

		if ok := scanner.Scan(); !ok {
			fmt.Println("Ошиюка ввода")
			return
		}

		text := scanner.Text()

		fields := strings.Fields(text)

		if len(fields) == 0 {
			fmt.Println("вы ничено не ввели")
			return
		}
		// fmt.Println("text", text)

		// pp.Println("слова", fields)
		fmt.Println("Команда:", fields[0])

		cmd := fields[0]

		if cmd == "выйти" {
			fmt.Println("выход...")
			return
		}

		if cmd == "добавить" {
			str := ""
			for i := 1; i < len(fields); i++ {
				str += fields[i]

				if i != len(fields)-1 {
					str += " "
				}
			}

			fmt.Println("вы хотите добавить:", str)
			fmt.Println("")

		} else if cmd == "удалить" {
			str := ""
			for i := 1; i < len(fields); i++ {
				str += fields[i]

				if i != len(fields)-1 {
					str += " "
				}
			}
			fmt.Println("вы хотите удалить:", str)
			fmt.Println("")

		} else if cmd == "help" {
			fmt.Println("команда: добавить {что нужно добавить}")
			fmt.Println("команда: удалить {что нужно удалить}")
			fmt.Println("команда: exit {завершить программу}")
			fmt.Println("команда: help {список команд}")
			fmt.Println("выйти {выход}")
			fmt.Println("")

		} else {
			fmt.Println("вы ввели неизвестную команду")
			fmt.Println("")
		}

	}

}
