package main

import (
	"errors"
	"fmt"
)

func Divide(number int, division int) (int, error) {
	if division == 0 {
		return 0, errors.New("Tidak bisa dibagi dengan 0")
	} else {
		return number / division, nil
	}
}

func main() {
	res, err := Divide(100, 0)

	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println("Hasil", res)
	}
}