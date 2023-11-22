package main

import "fmt"

func getFullName() (firstName string, middleName string, lastName string) {
	firstName = "Muhtarom"
	middleName = "Zain"
	lastName = "Muhammad Sondakh"

	return
}

func main() {
	a, b, c := getFullName()
	fmt.Println(a, b, c)
}