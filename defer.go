package main 

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApps() {
	defer logging() // walaupun ada error nantinya di akhir function akan tetap dipanggil
	fmt.Println("Run application")
	
	// error
	// error
	// error
	// error

	// cara normal letakkan di akhir, const jika terjadi error sebelum function logging func akan exit terlebih dahulu sebelum menjalankan logging
	// logging()
}

func main() {
	runApps()
}