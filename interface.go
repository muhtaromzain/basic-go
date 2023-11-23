package main

import "fmt"

type Person struct {
	Name, Address string
	Age int
}

type Animal struct {
	Name string
}

type HasName interface {
	GetName() string
}

func sayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{"Marc Marquez", "Madrid", 35}
	sayHello(person)

	animal := Animal{"Dog"}
	sayHello(animal)
}