package main

import "fmt"

func main() {
	// 9.1
	var a = 10
	var b = 20
	var c = 30
	var d = 40
	var e = 50

	var x = a + b - c * d / e
	fmt.Println(x)
	fmt.Println(x % x)
	
	// 9.2
	a += 10
	b -= 10
	c *= 10
	d /= 10
	e %= 10

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	
	// 9.3
	a++
	fmt.Println(a)
	a--
	fmt.Println(a)
}