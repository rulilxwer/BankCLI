package main

import (
	"bankCLI/pkg"
	"fmt"
)

func main() {
	var choice int
	fmt.Println("1. Создать клиента")
	fmt.Println("2. Показать баланс клиента")
	fmt.Println("3. Посмотреть баланс клиента")
	fmt.Println("4. Снять деньги c баланса клиента")
	fmt.Println("5. Выход")

	fmt.Scan(&choice)

	switch choice {
	case 1:
		pkg.RegisterClient()
	case 2:
		pkg.TopUpClientsAccount()
	case 3:
		pkg.ShowClientsBalance()
	case 4:
		pkg.WithdrawClientsAccount()
	case 5:
		return
	}
}
