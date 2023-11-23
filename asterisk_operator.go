package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	fmt.Println("------------")
	fmt.Println("WITHOUT ASTERISK")

	var address Address = Address{"Mataram", "Nusa Tenggara Barat", "Indonesia"}
	var address2 *Address = &address // reference to address (passing by reference)

	address2.City = "Selong"
	
	fmt.Println(address)
	fmt.Println(address2)
	
	fmt.Println("------------")

	address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"} // mengubah pointer pada variable pointer

	fmt.Println(address)
	fmt.Println(address2) // address2 memiliki data berbeda dengan address
	
	fmt.Println("------------")
	fmt.Println("WITH ASTERISK")

	var address3 Address = Address{"Mataram", "Nusa Tenggara Barat", "Indonesia"}
	var address4 *Address = &address3
	*address4 = Address{"Jakarta", "DKI Jakarta", "Indonesia"} // mengubah semua data pointer 
	
	fmt.Println(address3) // address3 memiliki data yang sama seperti address4
	fmt.Println(address4)

}