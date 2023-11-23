package main

import (
	"basic-go/database"
	_ "basic-go/internal" // this is blank identifier
	"fmt"
)

func main() {
	fmt.Println(database.GetDatabase())
}