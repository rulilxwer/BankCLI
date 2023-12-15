package pkg

import "fmt"

func TopUpClientsAccount() {
	var name string
	var amount int

	fmt.Println("")

	fmt.Scan(&name)

	balance, ok := database[name]

	if ok {
		fmt.Println("Ошибка!, данного пользователя нет в нашей бд")
	}

	fmt.Println("Введите сумму которую хотите пополнить")
	fmt.Scan(&amount)

	database[name] = balance + amount

}
