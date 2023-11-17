package main

import "fmt"

func main() {
	var name1 = "Zain"
	var name2 = "Zain"

	var result bool = name1 == name2
	var result2 bool = name1 != name2

	fmt.Println(result)
	fmt.Println(result2)
}