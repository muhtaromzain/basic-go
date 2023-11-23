package main

import "fmt"

type Man struct {
	Name string
}

func (man Man) Married() {
	man.Name = "Mr. " + man.Name
}

func (man *Man) MarriedUsingPointer() {
	man.Name = "Mr. " + man.Name
}

func main() {
	fmt.Println("------------")
	fmt.Println("WITHOUT POINTER METHOD")

	person := Man{"Zain"}
	person.Married()

	fmt.Println(person)

	fmt.Println("------------")
	fmt.Println("WITH POINTER METHOD")

	person.MarriedUsingPointer()

	fmt.Println(person)
}