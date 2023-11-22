package main

import "fmt"

func endApp() {
	fmt.Println("End Apps")
}

func runApps(error bool) {
	defer endApp()
	if error {
		panic("ERROR")
	}
}

func main() {
	runApps(true)
}