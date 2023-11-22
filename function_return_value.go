package main

import "fmt"

func getHello(name string)string {
	value := "Hello " + name
	return value
}

func main() {
	res := getHello("Zain")
	fmt.Println(res)
}