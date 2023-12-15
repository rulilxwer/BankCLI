package pkg

import "fmt"

func WithdrawClientsAccount() {
	var name string
	var amount int

	fmt.Println("Введите имя клиента")

	fmt.Scan(&name)

	balance, ok := database[name]

	if ok {
		fmt.Println("Ошибка!, данного пользователя нет в нашей бд")
	}

	fmt.Println("Введите сумму которую хотите взять")
	fmt.Scan(&amount)
	database[name] = balance - amount

	if balance < amount {
		fmt.Println("Недостаточно средств на балансе")
	}
}
