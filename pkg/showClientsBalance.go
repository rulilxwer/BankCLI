package pkg

import "fmt"

func ShowClientsBalance() {
	var name string

	fmt.Println("Введите имя клиента")

	fmt.Scan(&name)

	balance, ok := database[name]

	if ok {
		fmt.Println("Ошибка!, данного пользователя нет в нашей бд")
		return
	}
	fmt.Println("Баланс счета %n является %n", name, balance)
}
