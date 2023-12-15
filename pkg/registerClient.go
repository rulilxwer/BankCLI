package pkg

import "fmt"

func RegisterClient() {
	var name string
	var balance int

	fmt.Println("Введите имя клиента")

	fmt.Scan(&name)

	database[name] = balance

}
