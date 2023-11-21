package main

import "fmt"

func main() {
	name := "Muhtarom Zain"

	if name == "Zain" {
		fmt.Println("Hello", name)
	} else if name == "Muhtarom Zain" {
		fmt.Println("Hello Muhtarom")
	} else {
		fmt.Println("Hello guys")
	}

	if lenght := len(name); lenght > 5 {
		fmt.Println("Name is too long")
	} else {
		fmt.Println("Name OK")
	}
}