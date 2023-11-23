package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesiaWithoutPointer(address Address) {
	address.Country = "Indonesia"
}

func ChangeCountryToIndonesiaWithPointer(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	fmt.Println("------------")
	fmt.Println("WITHOUT POINTER FUNCTION")

	var address Address = Address{}
	ChangeCountryToIndonesiaWithoutPointer(address)

	fmt.Println(address)
	
	// jika terlanjur bukan pointer
	ChangeCountryToIndonesiaWithPointer(&address)

	fmt.Println(address)

	fmt.Println("------------")
	fmt.Println("WITH POINTER FUNCTION")

	var address2 *Address = &Address{}
	ChangeCountryToIndonesiaWithPointer(address2)

	fmt.Println(address2)
}