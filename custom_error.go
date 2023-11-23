package main

import "fmt"

type validationError struct {
	Message string
}

type notFoundError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

func (v *notFoundError) Error() string {
	return v.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{Message: "validation  error"}
	}

	if id != "Zain" {
		return &notFoundError{Message: "data not found"}
	}

	return nil
}

func main() {
	err := SaveData("Zain", nil)

	if err != nil {
		// using if
		// if validationError, ok := err.(*validationError); ok {
		// 	fmt.Println("validation error:", validationError.Error())
		// } else if notFoundError, ok := err.(*notFoundError); ok {
		// 	fmt.Println("not found error:", notFoundError.Error())
		// } else {
		// 	fmt.Println("unknown error", err.Error())
		// }

		// using switch

		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("validation error:", finalError.Error())
		case *notFoundError:
			fmt.Println("not found error:", finalError.Error())
		default:
			fmt.Println("unknown error", err.Error())
		}
	} else {
		fmt.Println("Berhasil")
	}
}