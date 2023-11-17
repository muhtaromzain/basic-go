package main

import "fmt"

func main() {
	// 5.1
	var name string

	name = "Muhtarom Zain"
	fmt.Println(name)

	name = "Zain Muhtarom"
	fmt.Println(name)

	var surname = "Muhammad Sondakh"
	fmt.Println(surname)

	lastname := "Sondakh"
	fmt.Println(lastname)

	lastname = "Sondakh Muhammad"
	fmt.Println(lastname)

	// 5.2
	var (
		firstName = "Muhtarom"
		middleName = "Zain"
		lastName = "Muhammad Sondakh"
	)

	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}