package main

import "fmt"

func sumAll(numbers ...int) (total int) {

	for _, number := range numbers {
		total += number
	}

	return
}

func sumAllUsingSlice(numbers []int) (total int) {

	for _, number := range numbers {
		total += number
	}

	return
}

func main() {
	res := sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(res)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	res2 := sumAllUsingSlice(numbers)
	fmt.Println(res2)

	// convert slice for variadic function
	res3 := sumAll(numbers...)
	fmt.Println(res3)
}