package main

import "fmt"

func factorial(number int) int {
	if number == 1 {
		return 1
	} else {
		return number * factorial(number - 1)
	}
}

func factorialLoop(value int) int {
	result := 1
	for x := value; x > 0; x-- {
		result *= x
	}

	return result
}

func main() {
	factorialLoop := factorialLoop(10)
	fmt.Println(factorialLoop)

	fmt.Println("------------")

	factorial := factorial(10)
	fmt.Println(factorial)
}