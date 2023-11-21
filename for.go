package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("counter:", counter)
		counter++
	}

	fmt.Println("Done")

	fmt.Println("--------------")

	for count := 1; count <= 10; count++ {
		fmt.Println("count:", count)
	}

	fmt.Println("Done")

	fmt.Println("--------------")

	names := []string{"Muhtarom", "Zain", "Muhammad", "Sondakh"}

	// manual
	for x := 0; x < len(names); x++ {
		fmt.Println(names[x])
	}

	fmt.Println("Done")

	fmt.Println("--------------")

	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}

	fmt.Println("Done")

	fmt.Println("--------------")

	// without index / key
	persons := map[string]string {
		"name": "Muhtarom",
		"address": "Bekasi",
	}

	for _, person := range persons {
		fmt.Println(person)
	}
}