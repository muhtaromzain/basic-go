package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "Muhtarom Zain",	
	}

	person["address"] = "Depok"
	person["address"] = "Bekasi" // overwrite data
	
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["age"]) // will return default value
	
	delete(person, "address")
	fmt.Println(person)
}