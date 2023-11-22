package main

import "fmt"

func endApp() {
	fmt.Println("End Apps")
}

func endAppWithRecover() {
	fmt.Println("End Apps")
	message := recover()
	fmt.Println("Terjadi panic", message)
}

func runAppsWithWrongRecover(error bool) {
	defer endApp()
	if error {
		panic("ERROR")
	}

	message := recover()
	fmt.Println("Terjadi panic", message)
}

func runAppsWithRightRecover(error bool) {
	defer endAppWithRecover()
	if error {
		panic("ERROR")
	}
}

func main() {
	// runAppsWithWrongRecover(true)
	runAppsWithRightRecover(true)
	fmt.Println("End")
}