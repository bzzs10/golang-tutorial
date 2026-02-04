package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func HelpConsole() {
	fmt.Println("Команда: добавить {инградиент} {количество} ---> добавляет что либо")
	fmt.Println("Команда: удалить {инградиент} {количество} ---> удаляет что либо")
	fmt.Println("Команда: получить {инградиент} ---> получение чего либо")
	fmt.Println("Команда: help ---> получение списка доступных команд")
	fmt.Println("Команда: выйти ---> выход из программы")
	fmt.Println("")
}

func main() {

	ingridient := make(map[string]int)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("введите команду: ")

		ok := scanner.Scan()
		if !ok {
			fmt.Println("ошибка ввода!")
			return
		}

		text := scanner.Text()

		fields := strings.Fields(text)
		if len(fields) == 0 {
			continue
		}

		cmd := fields[0]

		if cmd == "выйти" {
			fmt.Println("выход...")
			return
		}

		if cmd == "help" {
			HelpConsole()
			continue
		}

		if cmd == "добавить" {
			if len(fields) != 3 {
				fmt.Println("неверный формат команды")
				continue
			}

			amount, err := strconv.Atoi(fields[2])
			if err != nil {
				fmt.Println("Неверный формат количества ингридиента")
				continue
			}

			ingridient[fields[1]] += amount
			continue
		}

		if cmd == "удалить" {
			if len(fields) != 3 {
				fmt.Println("команда иммет неверный формат")
				continue
			}

			amount, err := strconv.Atoi(fields[2])
			if err != nil {
				fmt.Println("неверный формат количества ингридиентов")
				continue
			}

			ingridientAmount, ok := ingridient[fields[1]]
			if !ok {
				fmt.Println("попытка удалить несушествуюший ингридиент")
				continue
			}

			if ingridientAmount-amount < 0 {
				fmt.Println("попытка удалить больше ингриндиентов, чем было")
				continue
			}

			ingridient[fields[1]] = ingridientAmount - amount
		}
		if cmd == "получить" {
			if len(fields) != 2 {
				fmt.Println("неверный формат команды")
				continue
			}
			ingridientAmount, ok := ingridient[fields[1]]
			if !ok {
				fmt.Println("данный ингридиент не найден")
				continue
			}
			fmt.Println("Ингридиент:", fields[1], "количество:", ingridientAmount)
			continue
		}
		fmt.Println("передана неизвестная команда")
	}

}
