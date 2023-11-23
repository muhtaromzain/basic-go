package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	fmt.Println("------------")
	fmt.Println("EXAMPLE OF PASS BY VALUE")

	address1 := Address{"Depok", "Jawa Barat", "Indonesia"}
	address2 := address1 // copy value (passing by value)

	fmt.Println(address1)

	address2.City = "Bogor" // only address2 will change

	fmt.Println(address2)

	fmt.Println("------------")

	fmt.Println("EXAMPLE OF PASS BY REFERENCE") // using pointer

	address3 := Address{"Mataram", "Nusa Tenggara Barat", "Indonesia"}
	address4 := &address3 // reference to address3 (passing by reference)

	// var address3  Address = Address{"Mataram", "Nusa Tenggara Barat", "Indonesia"}
	// var address4 *Address = &address3 // reference to address3 (passing by reference)

	address4.City = "Selong"

	fmt.Println(address3)
	fmt.Println(address4)
	
	fmt.Println("------------")
}