package main

import "fmt"

func test() any { // type declaration for interface{}
	return 1
}

func main() {
	var kosong any = test()
	fmt.Println(kosong)
}