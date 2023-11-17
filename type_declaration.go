package main

import "fmt"

func main() {

	type nationalId string

	var NoKtp nationalId = "123456789"

	var contoh string = "111111111"
	var contohKtp nationalId = nationalId(contoh)

	fmt.Println(NoKtp)
	fmt.Println(contoh)
	fmt.Println(contohKtp)
}