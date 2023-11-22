package main

import "fmt"

type Filter func(string) (string)

func sayHelloWithFilter(name string, filter func(string) (string)) string {
	filteredName := filter(name)
	return filteredName
}

func sayHelloWithFilterWithTypeDeclaration(name string, filter Filter) string {
	filteredName := filter(name)
	return filteredName
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	res := sayHelloWithFilter("Anjing", spamFilter)
	fmt.Println(res)
	res2 := sayHelloWithFilterWithTypeDeclaration("Zain", spamFilter)
	fmt.Println(res2)
}