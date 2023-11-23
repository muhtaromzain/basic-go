package main

import (
	"fmt"
	"reflect"
)

func random() any {
	return true
}

func main() {
	var data any = random()
	var resultString string
	var resultInt int
	var resultBool bool

	fmt.Println(reflect.TypeOf(data))

	fmt.Println("----------")

	switch value := data.(type) {
	case string:
		fmt.Println("Data type:", value)
		resultString = data.(string)
		fmt.Println(reflect.TypeOf(resultString))
	case int:
		fmt.Println("Data type:", value)
		resultInt = data.(int)
		fmt.Println(reflect.TypeOf(resultInt))
	case bool:
		fmt.Println("Data type:", value)
		resultBool = data.(bool)
		fmt.Println(reflect.TypeOf(resultBool))
	default:
		fmt.Println("Data type:", value)
		fmt.Println("Not listed")
		break
	}

	fmt.Println("----------")
}