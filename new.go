package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	fmt.Println("------------")

	// var address *Address = &Address{}
	var address *Address = new(Address) // using new
	var address2 *Address = address
	
	address2.Country = "Indonesia"

	fmt.Println(address)
	fmt.Println(address2)
}