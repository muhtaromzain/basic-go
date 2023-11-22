package main

import "fmt"

func getFullName() (string, string) {
	return "Muhtarom", "Zain"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	_, sureName := getFullName()
	fmt.Println(sureName)
}