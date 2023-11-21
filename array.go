package main

import "fmt"

func main() {
	var names [4]string
	names[0] = "Muhtarom"
	names[1] = "Zain"
	names[2] = "Muhammad"
	names[3] = "Sondakh"

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])

	var values = [4]int{
		90,
		80,
		98,
	}

	fmt.Println(values)
	fmt.Println(values[1])
	fmt.Println(values[2])
	fmt.Println(values[3])
	fmt.Println(len(values))

	var price = [...]float32{
		16500.50,
		25000.50,
	}

	fmt.Println(price)
	fmt.Println(price[1])
	fmt.Println(len(price))
	price[len(price) + 1] = 2500.80
	fmt.Println(price)
	fmt.Println(len(price))
}