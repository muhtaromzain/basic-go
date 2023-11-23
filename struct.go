package main

import "fmt"

// 32.
type Person struct {
	Name, Address string
	Age int
}

// 33.
// struct method
func (person Person) sayHello(name string) {
	fmt.Println("Hello,", name, "my name is", person.Name)
}

func main() {
	var person Person
	person.Name = "Muhtarom Zain"
	person.Address = "Bekasi"
	person.Age = 25

	fmt.Println(person)

	person2 := Person{
		Name: "Joko Widodo",
		Address: "Tebet",
		Age: 40,
	}

	fmt.Println(person2)
	
	person3 := Person{"Marc Marquez", "Madrid", 35}
	
	fmt.Println(person3)
	// memanggil struct method
	person3.sayHello("Rossi")
}