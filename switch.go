package main

import "fmt"

func main() {
	name := "Zain"

	switch name {
	case "Muhtarom":
		fmt.Println("Hello Muhtarom")
	case "Zain":
		fmt.Println("Hello Muhtarom")
	default:
		fmt.Println("Hello, What is your name?")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Name is too long")
	case false:
		fmt.Println("Name OK")
	}

	// not recomended better use if
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Name is too long")
	case length > 4:
		fmt.Println("Name seems OK")
	default:
		fmt.Println("Name OK")
	}
}