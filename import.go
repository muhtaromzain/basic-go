package main

import (
	"basic-go/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("Zain")
	fmt.Println(result)\

	fmt.Println(helper.Application) // public
	fmt.Println(helper.version) // private
	fmt.Println(helper.sayGoodBye) // private
}