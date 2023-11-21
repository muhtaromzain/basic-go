package main

import "fmt"

func main() {
	for x := 0; x <= 10; x++ {
		if x == 5 {
			break
		}

		fmt.Println("loop:", x)
	}
}