package main

import "fmt"

func main() {
	for x := 0; x <= 10; x++ {
		if x%2 == 0 {
			continue
		}

		fmt.Println("loop:", x)
	}
}